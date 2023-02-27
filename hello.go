package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
 
type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"Isbn"`
	title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"id"`
	Lastnames string `json"Lastnames"`
}

var movies[]Movie
// var movies = []Movie{};  
// movies := make(map[Movie]float64)


func getMovies(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movies); 
}	
func deleteMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json");
	params  := mux.Vars(r);
	for index, item:= range movies{
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...);
			break;
		}
	}
	json.NewEncoder(w).Encode(movies)  ;
}
func getMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json");
	params:= mux.Vars(r);
	for _,item:= range movies{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item);
			break;
		}
	}
}
func createMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json");
	var movie Movie;
	json.NewDecoder(r.Body).Decode(&movie);
	var num int= rand.Intn(10000);
	movie.ID= strconv.Itoa(num);
	movies=append(movies, movie);
	json.NewEncoder(w).Encode(movies);
}
func updateMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type", "application/json");
	params:= mux.Vars(r);
	for index, items:= range movies{
		if items.ID == params["id"] {
			movies=append(movies[:index], movies[index+1:]...);
			var movie Movie;
			json.NewDecoder(r.Body).Decode(&movie);
			movie.ID= params["id"];
			movies=append(movies, movie); 
		}
	}
}
func main(){
	r:= mux.NewRouter();

	movies = append(movies, Movie{ID: "1", Isbn: "438", title: "movie 1", Director: &Director{Firstname: "john", Lastnames: "doe"}});
	movies = append(movies, Movie{ID: "2", Isbn: "439", title: "movie 2", Director: &Director{Firstname: "john", Lastnames: "doe"}});
	movies = append(movies, Movie{ID: "3", Isbn: "440", title: "movie 3", Director: &Director{Firstname: "john", Lastnames: "donut"}});


	r.HandleFunc("/movies", getMovies).Methods("GET");
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET");
	r.HandleFunc("/movies", createMovie).Methods("POST");
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT");
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE");

	fmt.Println("starting server at port 8000 \n");
	err := http.ListenAndServe(":8088", r)
	if err != nil {
			log.Fatalln(err)
	}



	
}