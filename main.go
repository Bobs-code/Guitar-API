package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Guitar struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Brand_id    int    `json:"brand_id"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

var Guitars []Guitar

func getGuitar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("Get single guitar endpoint hit")
	json.NewEncoder(w).Encode(Guitars)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/guitar", getGuitar)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Guitars = append(Guitars, Guitar{
		Id:          1,
		Name:        "Guitar Name",
		Brand_id:    1,
		Year:        3035,
		Description: "This is a description",
	})
	handleRequests()
}
