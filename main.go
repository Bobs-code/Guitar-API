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

	// Routes for Guitars resource
	r.Route("/guitars", func(r chi.Router) {
		r.Get("/", GetAllGuitars)
		r.Get("/{guitarId}", GetSingleGuitar)
		r.Put("/guitar/create", NewGuitar)
		r.Patch("/guitar/update", UpdateGuitar)
		r.Delete("/guitar/delete/{guitarId}", DeleteGuitar)
	})

	// Routes for Mucisian Resources
	r.Route("/mucisians", func(r chi.Router) {
		r.Get("/", AllMusicians)
		r.Get("/{musicianId}", Musician)
		r.Put("/musicians/create", NewMusician)
		r.Patch("/musicians/update/{musicianId}", UpdateMusician)
		r.Delete("/musician/delete{musicianId}", DeleteMusician)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequests()
}
