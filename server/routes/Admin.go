package routes

import (
	"log"
	"net/http"
)

// this is gonna be fun i can already feel it

func adminLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit AdminLogin")
	w.WriteHeader(200)
}
func deactivateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit DeactivateUser")
	w.WriteHeader(200)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit DeleteUser")
	w.WriteHeader(200)
}
