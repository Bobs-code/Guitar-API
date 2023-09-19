package conns

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	// Replace constants with correct values
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "placedholderc"
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
