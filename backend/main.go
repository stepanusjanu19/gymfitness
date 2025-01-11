package main

import (
	"log"
	"net/http"
	"backend/src"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("...............API GYM...............")
	log.Println(".....................................")
}

func main() {
	http.HandleFunc("/", Handler)
	src.Run()
}
