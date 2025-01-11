package handler

import (
	"log"
	"net/http"
	"api/src"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("...............API GYM...............")
	src.Run()
	log.Println(".....................................")
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    Handler(w, r)
}
