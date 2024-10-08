package main

import (
	authcontroller "fb-service/controller/authcontroller"
	"fb-service/controller/productcontroller"
	"fb-service/models"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.GetProducts).Methods("GET")
	api.HandleFunc("/product/{id}", productcontroller.GetProductById).Methods("GET")
	api.HandleFunc("/product", productcontroller.CreateProduct).Methods("POST")
	api.HandleFunc("/product/{id}", productcontroller.UpdateProduct).Methods("PUT")
	api.HandleFunc("/product/{id}", productcontroller.DeleteProduct).Methods("DELETE")
	// api.Use(middleware.JWTMiddleware)

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	corsHandler := corsOptions(r)

	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
