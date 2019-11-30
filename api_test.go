package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDefaultEvents(t *testing.T) {
	req, err := http.NewRequest("GET", "/events", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllEvents)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedValue := "Docker in Action"
	if !strings.Contains(rr.Body.String(), expectedValue) {
		t.Errorf("The expected value [%v] is not in the default json", expectedValue)
	}
}

func TestGetOneEvent(t *testing.T)  {
	req, err := http.NewRequest("GET", "/event/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/event/{id}", getOneEvent)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedValue := "Introduction to Golang"
	if !strings.Contains(rr.Body.String(), expectedValue) {
		t.Errorf("The expected value [%v] is not in the result json", expectedValue)
	}
}
