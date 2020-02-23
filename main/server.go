package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	userHandler := createUserHandler()

	router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/logout", userHandler.Logout)
	router.HandleFunc("/signup", userHandler.Add)
	router.HandleFunc("/user/{id:[0-9]+}", userHandler.UserById)

	log.Println("start serving :8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal("Error occurred while starting server", err)
	}
}
