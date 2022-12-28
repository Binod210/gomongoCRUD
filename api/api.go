package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Binod210/gomongoCRUD/db"
	"github.com/Binod210/gomongoCRUD/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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
	var users []model.User

	cursor, err := api.UserCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		json.NewEncoder(w).Encode("error in finding element")
		return
	}
	for cursor.Next(context.TODO()) {
		user := model.User{}
		err = cursor.Decode(&user)
		if err != nil {
			json.NewEncoder(w).Encode("Cannot Unmarsal element")
			continue
		}
		users = append(users, user)

	}
	json.NewEncoder(w).Encode(users)
	log.Println("GetAllUsers functin invoked")
}

func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateUser function invoked")
	params := mux.Vars(r)
	id := params["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode("cannot decode message")
		return
	}
	updateData, err := bson.Marshal(user)
	if err != nil {
		json.NewEncoder(w).Encode("cannot decode message")
		return
	}
	var Updat bson.M
	err = bson.Unmarshal(updateData, &Updat)
	update := bson.D{{Key: "$set", Value: Updat}}
	log.Println("data formed", update, objectId)
	result, err := api.UserCollection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)

	if err != nil {
		log.Println("error ", err)
		json.NewEncoder(w).Encode("Cannot Update user")
		return
	}
	if result.MatchedCount == 0 {
		json.NewEncoder(w).Encode("Cannot find user to update")
		return
	}
	json.NewEncoder(w).Encode("Successfully updated element")
	log.Println("Result ", result)

}
func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete User function invoked")
	params := mux.Vars(r)
	id := params["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := api.UserCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		json.NewEncoder(w).Encode("Cannot delete user")
		return
	}
	if result.DeletedCount == 0 {
		json.NewEncoder(w).Encode("User not found to delete")
		return
	}
	json.NewEncoder(w).Encode("Successfully deleted User")
	log.Printf("Id %v", id)

}
