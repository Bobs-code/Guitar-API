package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Initialize() *chi.Mux {

	r := chi.NewRouter()
	// Route for the home page
	// r.Get("/", homePage)

	// // Routes for Guitars resource
	// r.Route("/guitars", func(r chi.Router) {
	// 	r.Get("/", GetAllGuitars)
	// 	r.Get("/{guitarId}", GetSingleGuitar)
	// 	r.Put("/create", NewGuitar)
	// 	r.Patch("/update", UpdateGuitar)
	// 	r.Delete("/delete/{guitarId}", DeleteGuitar)
	// })

	// // Routes for Mucisian Resources
	// r.Route("/musicians", func(r chi.Router) {
	// 	r.Get("/", AllMusicians)
	// 	r.Get("/{musicianId}", Musician)
	// 	r.Put("/create", NewMusician)
	// 	r.Patch("/update/{musicianId}", UpdateMusician)
	// 	r.Delete("/delete/{musicianId}", DeleteMusician)
	// })

	return r
}

func ServerRouter() {
	r := Initialize()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error serving router")
	}
}
