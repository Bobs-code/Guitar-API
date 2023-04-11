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
	// Replace constants with correct values
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dbpasswordplaceholder"
	dbname   = "guitars"
)

type Guitar struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Brand_id    int    `json:"brand_id"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

type MultipleGuitars []Guitar

var multipleGuitars MultipleGuitars

func dbGetAllGuitars() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Open Postgres connection using above login statement
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// Query all Guitars from db
	sql := "select * FROM guitars "
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Error Query, and %s", err)
	}
	for rows.Next() {
		var eachGuitar Guitar
		err = rows.Scan(&eachGuitar.Id, &eachGuitar.Name, &eachGuitar.Brand_id, &eachGuitar.Year, &eachGuitar.Description)
		if err != nil {
			fmt.Printf("error Looping data, and %s", err)
		}
		multipleGuitars = append(multipleGuitars, eachGuitar)
	}
}

func getAllGuitars(w http.ResponseWriter, r *http.Request) {
	dbGetAllGuitars()
	w.Header().Set("Content-type", "application/json")
	fmt.Println("Get single guitar endpoint hit")
	json.NewEncoder(w).Encode(&multipleGuitars)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/guitar", getAllGuitars)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
