package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	//We need to convert all the fields into JSON
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "452329", Title: "Movie One",
		Director: &Director{FirstName: "Jhon", LastName: "Doe"}})

	movies = append(movies, Movie{ID: "2", Isbn: "45556329", Title: "Movie Two",
		Director: &Director{FirstName: "Adam", LastName: "NewMan"}})

	r := mux.NewRouter()
	//GET Request
	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovies).Methods("GET")
	//POST Request
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	//Delete Request
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Staring server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func getmovies(w http.ResponseWriter, r *http.Request) {
	//Defining conetent Type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func createMovie() {}
func updateMovie() {}
func deleteMovie() {}
