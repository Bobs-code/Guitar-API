package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetAllGuitars(w http.ResponseWriter, r *http.Request) {
	data := DbQueryAllGuitars()
	w.Header().Set("Content-type", "application/json")
	fmt.Println("All Guitars endpoint hit")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetSingleGuitar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := DbConnection()
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

// POST request INSERTING a guitar to the database
func NewGuitar(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()
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
func DeleteGuitar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db := DbConnection()

	// To retrieve a particular record form the database, we need to pass an id paremeter to the URL. We will use the following methods and assign it to the urlId variable
	// urlId := r.URL.Query().Get("id")
	urlId := chi.URLParam(r, "id")

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
func UpdateGuitar(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()
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
