package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var fetchBaconTextFunc = fetchBaconText

func fetchBaconText() (string, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func BeefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	text, err := fetchBaconTextFunc()
	if err != nil {
		http.Error(w, "Failed to fetch text", http.StatusInternalServerError)
		log.Println("Error:", err)

		return
	}

	meatCounts := CountMeats(text)
	response := map[string]map[string]int{"beef": meatCounts}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to fetch text", http.StatusInternalServerError)

		return
	}
}
