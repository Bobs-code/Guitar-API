package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func handleRequests() {
	r := chi.NewRouter()
	// Route for the home page
	r.Get("/", homePage)

	r.Route("/guitars", func(r chi.Router) {
		r.Get("/", GetAllGuitars)
		r.Get("/{guitarId}", GetSingleGuitar)
		r.Put("/guitar/create", NewGuitar)
		r.Patch("/guitar/update", UpdateGuitar)
		r.Delete("/guitar/delete/{guitarId}", DeleteGuitar)
	})

	// http.HandleFunc("/guitar", getSingleGuitar)
	// http.HandleFunc("/guitars", getAllGuitars)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequests()
}
