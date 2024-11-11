package main

import (
	"auth-service/routes"
	"log"
	"net/http"
)

func main() {
	mux := routes.SetupRoutes()


	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
