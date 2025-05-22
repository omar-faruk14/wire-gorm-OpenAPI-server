package main

import (
	"log"
	"net/http"
)

func main() {
	userHandler := InitUserHandler()

	http.HandleFunc("/users", userHandler.GetUsers)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
