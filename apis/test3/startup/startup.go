package startup

import (
	"fmt"
	"log"
	"main/apis/test3/services"
	"net/http"

	"github.com/gorilla/mux"
)

func Startup() {
	r := mux.NewRouter()

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
