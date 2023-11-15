package controller

import (
	"encoding/json"
	"fmt"
	"mongoApi/helpers"
	"mongoApi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := helpers.GetAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	inserted := helpers.InsertOneMovie(movie)
	fmt.Println(movie)

	movie.ID = inserted.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatchedController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	helpers.UpdateOneMovie(params["movieId"])
}

func DeleteOneMovieController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	deletedCount := helpers.DeleteOneMovie(params["movieId"])

	w.Write([]byte(fmt.Sprintf(`"deleted-count":%v`, deletedCount)))
}

func DeleteAllMovieController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	deletedCount := helpers.DeleteAllMovie()

	w.Write([]byte(fmt.Sprintf(`"deleted-count":%v`, deletedCount)))
}
