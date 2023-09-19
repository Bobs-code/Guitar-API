package controllers

import (
	"fmt"
	"log"

	"github.com/Bobs-code/Guitar-API/models"
	"github.com/jmoiron/sqlx"
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

var db *sqlx.DB

func InitPGDB() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to the databse: %v", err)
	}
	db = conn
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}

func GetGuitars() []models.Guitar {
	InitPGDB()
	defer db.Close()

	var guitars []models.Guitar
	// Query all Guitars from db
	sql := "SELECT * FROM guitars"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Error Query, and %s", err)
	}

	for rows.Next() {
		var eachGuitar models.Guitar
		err = rows.Scan(&eachGuitar.Id, &eachGuitar.Brand_id, &eachGuitar.Model, &eachGuitar.Year, &eachGuitar.Description)
		if err != nil {
			fmt.Printf("error Looping data, and %s", err)
		}
		guitars = append(guitars, eachGuitar)
	}
	return guitars
}
