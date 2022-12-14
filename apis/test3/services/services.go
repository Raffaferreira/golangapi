package services

import (
	"encoding/json"
	"fmt"
	obj "main/apis/test3/structs"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var movies []obj.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie obj.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie obj.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func HandleHello(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]string{
		"message": fmt.Sprintf("Hello %s", params["name"]),
	})
}

func HandlePerson(w http.ResponseWriter, rq *http.Request) {
	var p obj.Person
	err := json.NewDecoder(rq.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Nome: %s - Idade: %d\n", p.Name, p.Age)
}
