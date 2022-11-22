package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/praveenmahasena647/vue-link-tree/dbs"
	jwthandle "github.com/praveenmahasena647/vue-link-tree/jwt-handle"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	var data *dbs.AdminLogIn
	json.NewDecoder(r.Body).Decode(&data)
	if data == nil {
		w.WriteHeader(400)
		return
	}
	var admin *dbs.AdminLogIn
	var noEntry = dbs.AdminCollection.FindOne(context.TODO(), bson.M{"key": data.Key}).Decode(&admin)
	if noEntry != nil {
		w.WriteHeader(500)
		return
	}
	var passwordErr = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data.Password))
	if passwordErr != nil {
		w.WriteHeader(400)
		return
	}
	var tokenString, tokenErr = jwthandle.CreateAdminJWT(admin.Id.Hex())
	if tokenErr != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(tokenString))
}

func deactivateUser(w http.ResponseWriter, r *http.Request) {
	var id, tokenErr = jwthandle.DecodeAdminJWT(r.Header.Get("X-Admin"))
	if tokenErr != nil {
		w.WriteHeader(400)
		return
	}
	var dbId, idErr = primitive.ObjectIDFromHex(id)
	if idErr != nil {
		w.WriteHeader(400)
		return
	}

	var Exist = dbs.AdminCollection.FindOne(context.TODO(), bson.M{"_id": dbId})
	if Exist.Err().Error() != "" {
		w.WriteHeader(400)
		return
	}

	var data *dbs.AdminChange
	json.NewDecoder(r.Body).Decode(&data)
	var _, err = dbs.UserCollection.UpdateOne(context.TODO(), bson.M{"username": data.UserName}, bson.M{"$set": bson.M{"deactivate": true}})

	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var id, tokenErr = jwthandle.DecodeAdminJWT(r.Header.Get("X-Admin"))
	if tokenErr != nil {
		w.WriteHeader(400)
		return
	}

	var dbId, idErr = primitive.ObjectIDFromHex(id)
	if idErr != nil {
		w.WriteHeader(400)
		return
	}

	var Exist = dbs.AdminCollection.FindOne(context.TODO(), bson.M{"_id": dbId})
	if Exist.Err().Error() != "" {
		w.WriteHeader(400)
		return
	}

	var data *dbs.AdminChange
	json.NewDecoder(r.Body).Decode(&data)
	var _, err = dbs.UserCollection.DeleteOne(context.TODO(), bson.M{"username": data.UserName})
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)

}
