package server

import (
	"context"
	mongoentities "hse/link-accumulator/src/db/mongo-entities"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		w.Header().Add("status", "500")
	}

	user, err := mongoentities.GetUser(client, id)

	if err != nil {
		w.Header().Add("status", "500")
	}

	user_as_bytes, err := mongoentities.SerializeUser(&user)

	if err != nil {
		w.Header().Add("status", "500")
	}

	w.Write(user_as_bytes)
}
