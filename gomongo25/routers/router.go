package router

import (
	controller "github.com/Shiveshr140/gomongo/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

    router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.GetOneMovie).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controller.DeleteAllMovies).Methods("DELETE")

    return router
}