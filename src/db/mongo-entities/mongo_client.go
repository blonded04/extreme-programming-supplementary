package mongoentities

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName              = "BUDDY_DB"
	usersCollectionName = "USERS_COLLECTION"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateUser(client *mongo.Client, id string) error {
	collection := client.Database(dbName).Collection(usersCollectionName)

	newUser := User{
		id: id,
		root: Dir{
			subdirs: []Dir{},
			links: []Link{
				{link: "https://out_invite_link.com"},
			},
		},
	}

	_, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(client *mongo.Client, id string) (User, error) {
	collection := client.Database(dbName).Collection(usersCollectionName)

	var user User
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return User{}, err
	}

	if user.id == "" || len(user.root.name) != 0 {
		return User{}, fmt.Errorf("пользователь не найден или имеет недопустимые данные")
	}

	return user, nil
}

func UpdateUser(client *mongo.Client, id string, path string, link string) error {
	collection := client.Database(dbName).Collection(usersCollectionName)

	prev, err := GetUser(client, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	parts := strings.Split(path, "/")
	if parts[0] == "" {
		parts = []string{}
	}

	var updateDir func(dir *Dir, parts []string, link string) bool
	updateDir = func(dir *Dir, parts []string, link string) bool {
		if len(parts) == 0 {
			dir.links = append(dir.links, Link{link: link})
			return true
		}

		for i := range dir.subdirs {
			if dir.subdirs[i].name == parts[0] {
				return updateDir(&dir.subdirs[i], parts[1:], link)
			}
		}
		return false
	}

	if !updateDir(&prev.root, parts, link) {
		return errors.New("path not found")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"root": prev.root}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
