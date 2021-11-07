package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "AldiPC"
	DB_NAME     = "movies"
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disabled", DB_NAME, DB_PASSWORD, DB_USER)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return DB
}

type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies/", getMovies).Methods("GET")

	router.HandleFunc("/movies/", createMovie).Methods("POST")

	router.HandleFunc("/movies/{movieid}", deleteMovie).Methods("DELETE")

	router.HandleFunc("/movies/", deleteMovies).Methods("DELETE")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
