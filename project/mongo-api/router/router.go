package router

import (
	"mongoApi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovieController).Methods("POST")
	router.HandleFunc("/api/movie/{movieId}", controller.MarkAsWatchedController).Methods("PUT")
	router.HandleFunc("/api/movie/{movieId}", controller.DeleteOneMovieController).Methods("DELETE")
	router.HandleFunc("/api/delete-all-movies", controller.DeleteAllMovieController).Methods("DELETE")

	return router
}
