package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/praveenmahasena647/vue-link-tree/routes"
)

var handleCors = handlers.CORS(
	handlers.AllowedHeaders([]string{"Content-Type", "X-UserName", "X-UserEmail", "X-Jwt", "X-links"}),
	handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}),
	handlers.AllowedOrigins([]string{"http://localhost:5173"}),
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
		return
	}
	var router = routes.RunServer()
	http.ListenAndServe(":42069", handleCors(router))
}
