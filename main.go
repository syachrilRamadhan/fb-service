package main

import (
	authcontroller "fb-service/controller/authcontroller"
	"fb-service/controller/productcontroller"
	"fb-service/middleware"
	"fb-service/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middleware.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
