package main

import "net/http"

func AllMusicians(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Return All Musicians"))
}

func Musician(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Return Single"))
}

func NewMusician(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creates New Musician"))
}

func UpdateMusician(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updates Musician"))
}

func DeleteMusician(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletes Musician"))
}
