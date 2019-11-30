package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllEvents(t *testing.T) {
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

func TestGetOneEvent(t *testing.T) {
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

func TestCreateEvent(t *testing.T) {
	event := &event{
		ID:          "5",
		Title:       "Java8",
		Description: "Learn Java8 and Java Stream",
	}
	eventObject, _ := json.Marshal(event)
	req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(eventObject))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/event", createEvent)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedValue := "Java8"
	if !strings.Contains(rr.Body.String(), expectedValue) {
		t.Errorf("The expected value [%v] is not in the result json", expectedValue)
	}
}

func TestUpdateEvent(t *testing.T) {
	event := &event{
		ID:          "1",
		Title:       "Java8",
		Description: "Learn Java8 and Java Stream",
	}
	eventObject, _ := json.Marshal(event)
	req, err := http.NewRequest("POST", "/event/1", bytes.NewBuffer(eventObject))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/event/{id}", updateEvent)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedValue := "Java8"
	if !strings.Contains(rr.Body.String(), expectedValue) {
		t.Errorf("The expected value [%v] is not in the result json", expectedValue)
	}
}

func TestDeleteEvent(t *testing.T) {
	req, err := http.NewRequest("GET", "/event/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/event/{id}", deleteEvent)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedValue := "Introduction to Golang"
	if strings.Contains(rr.Body.String(), expectedValue) {
		t.Errorf("The expected value [%v] should not be in the result json", expectedValue)
	}
}
