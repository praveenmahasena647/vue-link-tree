package routes

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/praveenmahasena647/vue-link-tree/dbs"
	jwthandle "github.com/praveenmahasena647/vue-link-tree/jwt-handle"
	"go.mongodb.org/mongo-driver/bson"
)

// this is gonna be fun i can already feel it
// lets hate JSML together

func adminLogin(w http.ResponseWriter, r *http.Request) {
	var Admin *dbs.AdminLogIn = &dbs.AdminLogIn{}
	json.NewDecoder(r.Body).Decode(&Admin)
	var token, tokenErr = jwthandle.CreateAdminJWT(Admin.Key)
	if tokenErr != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(token))
}

func deactivateUser(w http.ResponseWriter, r *http.Request) {
	var data, dataErr = io.ReadAll(r.Body)
	if dataErr != nil {
		w.WriteHeader(500)
		return
	}
	var _, notDone = dbs.UserCollection.UpdateOne(context.TODO(), bson.M{"username": string(data)}, bson.M{"$set": bson.M{"deactivate": false}})
	if notDone != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var data, dataErr = io.ReadAll(r.Body)
	if dataErr != nil {
		w.WriteHeader(500)
		return
	}
	var done, notDelete = dbs.UserCollection.DeleteOne(context.TODO(), bson.M{"username": string(data)})
	log.Println(done, notDelete)

	w.WriteHeader(200)
}
