package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yo, World!")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /yo", hiHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
