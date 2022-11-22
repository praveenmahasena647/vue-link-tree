package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/praveenmahasena647/vue-link-tree/dbs"
	jwthandle "github.com/praveenmahasena647/vue-link-tree/jwt-handle"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	var data *dbs.AdminLogIn
	json.NewDecoder(r.Body).Decode(&data)
	if data == nil {
		w.WriteHeader(500)
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
		w.WriteHeader(500)
		return
	}
	var tokenString, tokenErr = jwthandle.CreateAdminJWT(admin.Id.Hex())
	if tokenErr != nil {

		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(tokenString))
}

func deactivateUser(w http.ResponseWriter, r *http.Request) {
	var tokenString string = r.Header.Get("X-Admin")
	log.Println(tokenString)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var tokenString string = r.Header.Get("X-Admin")
	log.Println(tokenString)
}
