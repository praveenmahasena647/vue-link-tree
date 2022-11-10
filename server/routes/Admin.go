package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/praveenmahasena647/vue-link-tree/dbs"
	jwthandle "github.com/praveenmahasena647/vue-link-tree/jwt-handle"
)

// this is gonna be fun i can already feel it

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
	var data = map[string]string{}

	json.NewDecoder(r.Body).Decode(&data)
	log.Println(data)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{}

	json.NewDecoder(r.Body).Decode(&data)
	log.Println(data)
}
