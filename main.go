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
	password = "!1005XOctoberX!2989"
	dbname   = "guitars"
)

type Guitar struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Brand_id    int    `json:"brand_id"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

// Make connection to the database
func dbConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Open Postgres connection using above login statement
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")
	return db
}

// Select all guitars from database
func dbReturnAllGuitars() []Guitar {
	db := dbConnection()
	var multipleGuitars []Guitar
	// Query all Guitars from db
	sql := "SELECT * FROM guitars "
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
	return multipleGuitars
}

// GET request for return data from dbReturnAllGuitars()
func getAllGuitars(w http.ResponseWriter, r *http.Request) {
	data := dbReturnAllGuitars()
	w.Header().Set("Content-type", "application/json")
	fmt.Println("Get single guitar endpoint hit")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// POST request adding a guitar to the database
func addGuitar(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()
	w.Header().Set("Content-type", "application/json")

	var guitar Guitar
	err := json.NewDecoder(r.Body).Decode(&guitar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `
	INSERT INTO guitars (name, brand_id, year, description)
	VALUES ($1, $2, $3, $4)
	returning id`
	id := 0
	err = db.QueryRow(sqlStatement, guitar.Name, guitar.Brand_id, guitar.Year, guitar.Description).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Item with ID %d was created", id)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/guitars", getAllGuitars)
	http.HandleFunc("/new", addGuitar)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
