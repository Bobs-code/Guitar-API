package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Home page!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
