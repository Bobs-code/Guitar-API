package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Musician struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Guitar int    `json:"primary_guitar_id"`
}

// Get access to all of the muisician resources
func dbQueryAllMusicians() []Musician {
	db := DbConnection()
	var multipleMusicians []Musician

	sql := "SELECT * FROM musicians"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Querry Error, and %s", err)

	}

	for rows.Next() {
		var Musician Musician
		err = rows.Scan(&Musician.Id, &Musician.Name, &Musician.Guitar)
		if err != nil {
			fmt.Printf("error Looping over data, and %s", err)

		}
		if Musician.Guitar == 1 {
			fmt.Println("Fends")
		}

		multipleMusicians = append(multipleMusicians, Musician)
	}

	return multipleMusicians
}

// Return all musician resources
func GetAllMusicians(w http.ResponseWriter, r *http.Request) {
	data := dbQueryAllMusicians()
	w.Header().Set("Content-type", "application/json")
	fmt.Println("Returning All Musicians")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Get a single musician resource
func GetMusician(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db := DbConnection()
	defer db.Close()

	urlId := r.URL.Query().Get("id")

	urlIdInt, err := strconv.Atoi(urlId)
	if err != nil {
		panic(err)
	}

	sqlStatement := "SELECT * FROM musicians WHERE ID = $1"

	row := db.QueryRow(sqlStatement, urlIdInt)

	var musician Musician

	switch err := row.Scan(&musician.Id, &musician.Name, &musician.Guitar); err {
	case sql.ErrNoRows:
		fmt.Println("nno rows were returned.")
	case nil:
		fmt.Println(`record from the database: `, musician)
	default:
		panic(err)
	}

	err = json.NewEncoder(w).Encode(musician)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Create a new musician resource
func AddMusician(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()
	defer db.Close()
	w.Header().Set("Content-type", "application/json")

	var musician Musician

	err := json.NewDecoder(r.Body).Decode(&musician)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sql := `INSERT INTO musicians (name, primary_guitar_id)
	VALUES ($1, $2)
	returning id`

	id := 0

	err = db.QueryRow(sql, musician.Name, musician.Guitar).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "Musician with Id %d was created", id)
}

//
