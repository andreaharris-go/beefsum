package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/beef/summary", BeefSummaryHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
