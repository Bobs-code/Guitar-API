package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Bobs-Code/Guitar-API/conns"

	"github.com/gin-gonic/gin"
)

type Guitar struct {
	Id          int    `json:"id"`
	Brand_id    int    `json:"brand_id"`
	Model       string `json:"model"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

// SELECT all guitars from database
func dbQueryAllGuitars() []Guitar {

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
func getSingleGuitar(c *gin.Context) {

	// To retrieve a particular record form the database, we need to pass an id paremeter to the URL. We will use the following methods and assign it to the urlId variable
	urlId := c.Param("id")

	// To add a layer of security, we will cast the urlId param to an integer from a string. This will be passed into the database query below.
	urlIdInt, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sqlStatement := "SELECT * FROM guitars WHERE id = $1;"

	row := db.QueryRow(sqlStatement, urlIdInt)

	var singleGuitar Guitar

	switch err := row.Scan(&singleGuitar.Id, &singleGuitar.Brand_id, &singleGuitar.Model, &singleGuitar.Year, &singleGuitar.Description); err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	case nil:
		fmt.Println("Record from the database: ", singleGuitar)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, singleGuitar)
}

// GET request to return data from dbReturnAllGuitars()
func getAllGuitars(c *gin.Context) {
	data := dbQueryAllGuitars()
	c.JSON(http.StatusOK, data)
	fmt.Println("Get all guitars endpoint hit")
}

// POST request INSERTING a guitar to the database
func newGuitar(c *gin.Context) {

	var guitar Guitar

	if err := c.ShouldBindJSON(&guitar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `
	INSERT INTO guitars (brand_id, model, year, description)
	VALUES ($1, $2, $3, $4)
	returning id`

	var id int
	err := db.QueryRow(sqlStatement, guitar.Brand_id, guitar.Model, guitar.Year, guitar.Description).Scan(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guitar.Id = id
	c.JSON(http.StatusCreated, guitar)
	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "Item with ID %d was created", id)
}

// DELETE request
func deleteGuitar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

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

}

// Update request
func updateGuitar(w http.ResponseWriter, r *http.Request) {

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

func homePage(c *gin.Context) {
	c.Writer.WriteString("GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func handleRequests() {

	http.HandleFunc("/guitar/update", updateGuitar)
	http.HandleFunc("/guitar/delete", deleteGuitar)
}

func main() {
	conns.InitPGDB()
	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/guitar/:id", getSingleGuitar)
	r.GET("/guitars", getAllGuitars)
	r.PUT("/guitar/create", newGuitar)

	r.Run("localhost:8080")

}
