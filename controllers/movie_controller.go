package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/ParthLukhi/Go-Kubernetes/models"
	"github.com/gorilla/mux"
)

func RegisterMovieRoutes(router *mux.Router) {
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	models.Movies = append(models.Movies, models.Movie{ID: "0", Isbn: "438777", Title: "Movie 0", Director: &models.Director{Firstname: "John", Lastname: "Wick"}})
	models.Movies = append(models.Movies, models.Movie{ID: "1", Isbn: "432897", Title: "Movie 1", Director: &models.Director{Firstname: "Ben", Lastname: "Dover"}})
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range models.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	models.Movies = append(models.Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Movies {
		if item.ID == params["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			models.Movies = append(models.Movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Movies {
		if item.ID == params["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Movies)
}
