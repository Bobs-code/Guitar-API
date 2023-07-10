package driver

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



type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf("root:%s$tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)

	d, err := sql.Open("psql", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
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
