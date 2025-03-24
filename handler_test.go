package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBeefSummaryHandler_Success(t *testing.T) {
	fetchBaconTextFunc = func() (string, error) {
		return "corned ham corned brisket", nil
	}

	defer func() { fetchBaconTextFunc = fetchBaconText }()

	req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
	rec := httptest.NewRecorder()

	BeefSummaryHandler(rec, req)
	res := rec.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", res.StatusCode)
	}

	var body map[string]map[string]int
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	beef, ok := body["beef"]
	if !ok {
		t.Fatalf("Missing 'beef' field in response")
	}

	if beef["corned"] != 2 || beef["ham"] != 1 || beef["brisket"] != 1 {
		t.Errorf("Unexpected count result: %v", beef)
	}
}

func TestBeefSummaryHandler_FetchFail(t *testing.T) {
	fetchBaconTextFunc = func() (string, error) {
		return "", errors.New("mock fetch error")
	}

	defer func() { fetchBaconTextFunc = fetchBaconText }()

	req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
	rec := httptest.NewRecorder()

	BeefSummaryHandler(rec, req)
	res := rec.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Expected 500 error, got %d", res.StatusCode)
	}
}
