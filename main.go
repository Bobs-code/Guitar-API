package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
	Brand_id    int    `json:"brand_id"`
	Model       string `json:"model"`
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

// SELECT all guitars from database
func dbQueryAllGuitars() []Guitar {
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
		err = rows.Scan(&eachGuitar.Id, &eachGuitar.Brand_id, &eachGuitar.Model, &eachGuitar.Year, &eachGuitar.Description)
		if err != nil {
			fmt.Printf("error Looping data, and %s", err)
		}
		multipleGuitars = append(multipleGuitars, eachGuitar)
	}
	return multipleGuitars
}

// GET guitar record form dbQuerySingleRecord
func getSingleGuitar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dbConnection()
	defer db.Close()
	// To retrieve a particular record form the database, we need to pass an id paremeter to the URL. We will use the following methods and assign it to the urlId variable
	// urlId := r.URL.Query().Get("id")
	urlId := chi.URLParam(r, "guitarId")

	// To add a layer of security, we will cast the urlId param to an integer from a string. This will be passed into the database query below.
	urlIdInt, err := strconv.Atoi(urlId)
	if err != nil {
		panic(err)
	}

	sqlStatement := "SELECT * FROM guitars WHERE id = $1;"

	row := db.QueryRow(sqlStatement, urlIdInt)

	var singleGuitar Guitar

	switch err := row.Scan(&singleGuitar.Id, &singleGuitar.Brand_id, &singleGuitar.Model, &singleGuitar.Year, &singleGuitar.Description); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(`Record from the database: `, singleGuitar)
	default:
		panic(err)
	}

	err = json.NewEncoder(w).Encode(singleGuitar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GET request to return data from dbReturnAllGuitars()
func getAllGuitars(w http.ResponseWriter, r *http.Request) {
	data := dbQueryAllGuitars()
	w.Header().Set("Content-type", "application/json")
	fmt.Println("Get single guitar endpoint hit")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// POST request INSERTING a guitar to the database
func newGuitar(w http.ResponseWriter, r *http.Request) {
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
	INSERT INTO guitars (brand_id, model, year, description)
	VALUES ($1, $2, $3, $4)
	returning id`
	id := 0
	err = db.QueryRow(sqlStatement, guitar.Brand_id, guitar.Model, guitar.Year, guitar.Description).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Item with ID %d was created", id)
}

// DELETE request
func deleteGuitar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db := dbConnection()

	// To retrieve a particular record form the database, we need to pass an id paremeter to the URL. We will use the following methods and assign it to the urlId variable
	urlId := r.URL.Query().Get("id")

	// To add a layer of security, we will cast the urlId param to an integer from a string. This will be passed into the database query below.
	urlIdInt, err := strconv.Atoi(urlId)
	if err != nil {
		panic(err)
	}

	sqlStatement := "DELETE FROM guitars WHERE id = $1;"

	res, err := db.Exec(sqlStatement, urlIdInt)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
	defer db.Close()
}

// Update request
func updateGuitar(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	urlId := r.URL.Query().Get("id")

	urlIdInt, err := strconv.Atoi(urlId)

	if err != nil {
		panic(err)
	}

	var guitar Guitar

	err = json.NewDecoder(r.Body).Decode(&guitar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `
	UPDATE guitars
	SET brand_id = $2, model = $3, year = $4, description = $5
	WHERE ID = $1;
	`
	_, err = db.Exec(sqlStatement, urlIdInt, guitar.Brand_id, guitar.Model, guitar.Year, guitar.Description)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusAccepted)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func handleRequests() {
	r := chi.NewRouter()
	r.Route("/guitars", func(r chi.Router) {
		r.Get("/", getAllGuitars)
		r.Get("/{guitarId}", getSingleGuitar)
	})
	http.HandleFunc("/", homePage)
	http.HandleFunc("/guitar/create", newGuitar)
	http.HandleFunc("/guitar/update", updateGuitar)
	http.HandleFunc("/guitar/delete", deleteGuitar)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequests()
}
