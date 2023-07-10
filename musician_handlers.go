package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Musician struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Guitar int    `json:"primary_guitar_id"`
}

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

	err = db.QueryRow(sql, musician.Id, musician.Name, musician.Guitar).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "Musician with Id %d was created", id)
}
