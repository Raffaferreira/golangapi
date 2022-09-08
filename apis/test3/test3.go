package main

import (
	"fmt"
	"log"
	services "main/apis/test3/services"
	obj "main/apis/test3/structs"
	"net/http"

	"github.com/gorilla/mux"
)

var movies []obj.Movie

func main() {

	r := mux.NewRouter()

	movies = append(movies, obj.Movie{Id: "1", Isbn: "434779", Title: "Movie One", Director: &obj.Director{FirstName: "John", LastName: "Mayer"}})
	movies = append(movies, obj.Movie{Id: "2", Isbn: "434780", Title: "Movie Two", Director: &obj.Director{FirstName: "Rafael", LastName: "Ferreira"}})

	r.HandleFunc("/{name}", services.HandleHello).Methods("GET")
	r.HandleFunc("/person", services.HandlePerson).Methods("POST")

	r.HandleFunc("/movies", services.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", services.GetMovieById).Methods("GET")
	r.HandleFunc("/create", services.CreateMovie).Methods("POST")
	r.HandleFunc("/update/{id}", services.UpdateMovie).Methods("PUT")
	r.HandleFunc("/delete/{id}", services.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
