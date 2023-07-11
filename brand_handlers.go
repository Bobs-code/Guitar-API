package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Brand struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SELECT all brands from database
func DbQueryAllbrands() []Brand {
	db := DbConnection()
	var multiplebrands []Brand
	// Query all brands from db
	sql := "SELECT * FROM brand"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Error Query, and %s", err)
	}

	for rows.Next() {
		var brand Brand
		err = rows.Scan(&brand.Id, &brand.Name, &brand.Description)
		if err != nil {
			fmt.Printf("error Looping data, and %s", err)
		}
		multiplebrands = append(multiplebrands, brand)
	}
	return multiplebrands
}

// GET brand record form dbQuerySingleRecord
func GetBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db := DbConnection()
	defer db.Close()
	// To retrieve a particular record form the database, we need to pass an id paremeter to the URL. We will use the following methods and assign it to the urlId variable
	urlId := r.URL.Query().Get("id")

	// To add a layer of security, we will cast the urlId param to an integer from a string. This will be passed into the database query below.
	urlIdInt, err := strconv.Atoi(urlId)
	if err != nil {
		panic(err)
	}

	sqlStatement := "SELECT * FROM brand WHERE id = $1;"

	row := db.QueryRow(sqlStatement, urlIdInt)

	var brand Brand

	switch err := row.Scan(&brand.Id, &brand.Name, &brand.Description); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(`Record from the database: `, brand)
	default:
		panic(err)
	}

	err = json.NewEncoder(w).Encode(brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GET request to return data from dbReturnAllbrands()
func GetAllBrands(w http.ResponseWriter, r *http.Request) {
	data := DbQueryAllbrands()
	w.Header().Set("Content-type", "application/json")
	fmt.Println("Get single brand endpoint hit")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// POST request INSERTING a brand to the database
func NewBrand(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()
	defer db.Close()
	w.Header().Set("Content-type", "application/json")

	var brand Brand
	err := json.NewDecoder(r.Body).Decode(&brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `
	INSERT INTO brand (name, description)
	VALUES ($1, $2)
	returning id`
	id := 0
	err = db.QueryRow(sqlStatement, brand.Name, brand.Description).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Item with ID %d was created", id)
}

// DELETE request
func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db := DbConnection()

	urlId := r.URL.Query().Get("id")

	urlIdInt, err := strconv.Atoi(urlId)
	if err != nil {
		panic(err)
	}

	sqlStatement := "DELETE FROM brand WHERE id = $1;"

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
func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	urlId := r.URL.Query().Get("id")

	urlIdInt, err := strconv.Atoi(urlId)

	if err != nil {
		panic(err)
	}

	var brand Brand

	err = json.NewDecoder(r.Body).Decode(&brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `
	UPDATE brand
	SET name = $2, description = $3
	WHERE ID = $1;
	`
	_, err = db.Exec(sqlStatement, urlIdInt, brand.Name, brand.Description)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusAccepted)
}
