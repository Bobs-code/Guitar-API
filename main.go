package main

import (
	"database/sql"
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
	password = "placeholder"
	dbname   = "guitars"
)

// Make connection to the database
func DbConnection() *sql.DB {
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func guitarRequests() {
	http.HandleFunc("/guitar", GetSingleGuitar)
	http.HandleFunc("/guitars", GetAllGuitars)
	http.HandleFunc("/guitar/create", NewGuitar)
	http.HandleFunc("/guitar/update", UpdateGuitar)
	http.HandleFunc("/guitar/delete", DeleteGuitar)
}

func musicianRequests() {
	http.HandleFunc("/musician", GetMusician)
	http.HandleFunc("/musicians", GetAllMusicians)
	http.HandleFunc("/musicians/create", AddMusician)
	http.HandleFunc("/musicians/update", UpdateMusician)
	http.HandleFunc("/musicians/delete", DeleteMusician)
}

func brandRequests() {
	http.HandleFunc("/brand", GetBrand)
	http.HandleFunc("/brands", GetAllBrands)
	http.HandleFunc("/brands/create", NewBrand)
	http.HandleFunc("/brands/update", UpdateBrand)
	http.HandleFunc("/brands/delete", DeleteBrand)
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	guitarRequests()
	musicianRequests()
	brandRequests()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
