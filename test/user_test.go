package main

import (
	"api/infra/controllers"
	"api/infra/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserListWhenIsOk(t *testing.T) {

	ctl := controllers.NewUserController(
		database.NewConnect("postgresql://postgres:postgres@localhost:5432/postgres"),
	)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetUsers())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"users":null}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserListWhenHasQuery(t *testing.T) {

	ctl := controllers.NewUserController(
		database.NewConnect("postgresql://postgres:postgres@localhost:5432/postgres"),
	)

	req, err := http.NewRequest("GET", "/users?test=11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctl.GetUsers())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"users":null}`

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
