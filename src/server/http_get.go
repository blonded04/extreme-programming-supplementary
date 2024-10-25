package server

import (
	"context"
	mongoentities "hse/link-accumulator/src/db/mongo-entities"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// @Summary Get user
// @Description Get user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} Users
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /api/get_user_data/{id} [get]
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
