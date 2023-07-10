package main

import (
	"fmt"
	"net/http"
)

func AllMusicians(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Return All Musicians"))
	if err != nil {
		fmt.Println(err)
	}
}

func Musician(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Return Single"))
	if err != nil {
		fmt.Println(err)
	}

}

func NewMusician(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Creates New Musician"))
	if err != nil {
		fmt.Println(err)
	}

}

func UpdateMusician(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Updates Musician"))
	if err != nil {
		fmt.Println(err)
	}

}

func DeleteMusician(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Deletes Musician"))
	if err != nil {
		fmt.Println(err)
	}

}
