package mongoentities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success create user", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := CreateUser(mt.Client, "test_id")

		assert.NoError(t, err)
	})
}

func TestGetUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.AddMockResponse(
		&mtest.MockResponse{
			Coll:       mt.Coll,
			Collection: "users",
			Filter:     bson.M{"_id": "test_id"},
			Response:   bson.Raw(`{"_id":"test_id","name":"Test User"}`),
		},
	)

	coll := mt.Coll

	mockClient := mongo.NewMockClient(mt.ClientOptions())

	mt.MockClient = mockClient

	mt.Run("success get user", func(mt *mtest.T) {
		user := User{id: "test_id"}

		firstResponse := mtest.CreateCursorResponse(1, "BUDDY_DB.USERS_COLLECTION", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: "test_id"},
			{Key: "root", Value: bson.D{
				{Key: "subdirs", Value: bson.A{}},
				{Key: "links", Value: bson.A{
					bson.D{{Key: "link", Value: "https://out_invite_link.com"}},
				}},
			}},
		})

		mt.AddMockResponses(firstResponse)

		result, err := GetUser(mt.MockClient, "test_id")

		assert.NoError(t, err)
		assert.Equal(t, user.id, result.id)
	})
}

func TestUpdateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.AddMockResponse(
		&mtest.MockResponse{
			Coll:       mt.Coll,
			Collection: "users",
			Filter:     bson.M{"_id": "test_id"},
			Response:   bson.Raw(`{"_id":"test_id","name":"Test User"}`),
		},
	)

	coll := mt.Coll

	mockClient := mongo.NewMockClient(mt.ClientOptions())

	mt.MockClient = mockClient

	mt.Run("success update user", func(mt *mtest.T) {
		firstResponse := mtest.CreateCursorResponse(1, "BUDDY_DB.USERS_COLLECTION", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: "test_id"},
			{Key: "root", Value: bson.D{
				{Key: "subdirs", Value: bson.A{}},
				{Key: "links", Value: bson.A{
					bson.D{{Key: "link", Value: "https://out_invite_link.com"}},
				}},
			}},
		})
		mt.AddMockResponses(firstResponse, mtest.CreateSuccessResponse())

		err := UpdateUser(mt.MockClient, "test_id", "", "https://new_link.com")
		assert.NoError(t, err)
	})
}
