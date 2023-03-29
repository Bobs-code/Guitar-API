package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "passwordplaceholder"
	dbname   = "guitars"
)

type Guitar struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Brand_id    int    `json:"brand_id"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

type Guitars []Guitar

func getGuitars(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// quick test
	guitars := Guitars{
		Guitar{Id: 1, Name: "Guitar Name"},
	}

	json.NewEncoder(writer).Encode(guitars)
}

func handleRequests() {
	http.HandleFunc("/api/v1/Guitars", getGuitars)
	fmt.Println("Endpoint Hit: All Guitars endpoint")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// handleRequests()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("No error and successfully connected")
}
