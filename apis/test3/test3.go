package main

import (
	startup "main/apis/test3/startup"
	obj "main/apis/test3/structs"
)

var movies []obj.Movie

func main() {

	movies = append(movies, obj.Movie{Id: "1", Isbn: "434779", Title: "Movie One", Director: &obj.Director{FirstName: "John", LastName: "Mayer"}})
	movies = append(movies, obj.Movie{Id: "2", Isbn: "434780", Title: "Movie Two", Director: &obj.Director{FirstName: "Rafael", LastName: "Ferreira"}})
	startup.Startup()
}
