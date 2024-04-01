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
	http.HandleFunc("/yo", hiHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
