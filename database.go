package main

import (
	"database/sql"
	"fmt"
)

const (
	// Replace constants with correct values
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "placeholder"
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

// SELECT all guitars from database
func DbQueryAllGuitars() []Guitar {
	db := DbConnection()
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
