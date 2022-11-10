package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func middleWare(h http.Handler) http.Handler {
	// gotta do a learn on it
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		h.ServeHTTP(w, r)
	})
}

func RunServer() *mux.Router {
	var router *mux.Router = mux.NewRouter()
	var User = router.PathPrefix("/User").Subrouter()
	var Admin = router.PathPrefix("/Admin").Subrouter()
	// get routes -- for visits
	User.HandleFunc("/", getAllUsers).Methods("GET")
	User.HandleFunc("/{name}", getOneUser).Methods("GET")
	User.HandleFunc("/verify/", getMail).Methods("GET")
	User.HandleFunc("/c/", checkForUsers).Methods("GET")
	// Create routes for users
	User.HandleFunc("/", createUser).Methods("POST")
	User.HandleFunc("/login/", viewDashBoard).Methods("GET")
	User.HandleFunc("/login", handleLogin).Methods("POST")
	User.HandleFunc("/links", updateLink).Methods("PUT")
	User.HandleFunc("/changePassword", changePassword).Methods("POST")
	User.HandleFunc("/active", active).Methods("POST")
	User.HandleFunc("/getVerified/{id}", verify).Methods("GET")

	Admin.HandleFunc("/", adminLogin).Methods("POST")
	Admin.HandleFunc("/deactivateUser", deactivateUser).Methods("POST")
	Admin.HandleFunc("/deleteUser", deleteUser).Methods("POST")

	return router
}
