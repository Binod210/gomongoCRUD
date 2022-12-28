package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Binod210/gomongoCRUD/db"
	"github.com/Binod210/gomongoCRUD/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type API struct {
	UserCollection *mongo.Collection
}

func CreateApi(mongoConnString string) *API {
	mongoClient, err := db.StartDb(mongoConnString)
	if err != nil {
		log.Fatalf("Couldnot connect to DB %v", err)
	}
	collection := mongoClient.Database("ContactDb").Collection("contacts")
	return &API{UserCollection: collection}
}

func (api *API) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Create User function invoked")
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode("cannot decode message")
		return
	}
	log.Println("Received message ", user)
	result, err := api.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		json.NewEncoder(w).Encode("error in saving user")
	}
	response := model.UserResponse{
		Id:    result.InsertedID.(primitive.ObjectID).Hex(),
		Name:  user.Name,
		Email: user.Email,
	}
	json.NewEncoder(w).Encode(response)

}

func (api *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllUsers functin invoked")
}

func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateUser function invoked")
}
func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
