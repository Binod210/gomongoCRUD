package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Binod210/gomongoCRUD/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type API struct {
	DB *mongo.Client
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
	json.NewEncoder(w).Encode(user)

}

func (api *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllUsers functin invoked")
}

func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateUser function invoked")
}
func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
