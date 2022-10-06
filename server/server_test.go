package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatal(err)
	}
	//server.DB.Connect()
	if server.DB.Queries == nil {
		t.Errorf("Server DB failed to initialize. Ensure the testing database is running..")
	}

}

func TestGetIndex(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// like a response writer
	w := httptest.NewRecorder()
	// make sure index page serves
	server.GetIndex(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	findhtml := regexp.MustCompile(`\s*<html`)
	is_html := findhtml.Match(data)
	// is_html := findhtml.MatchString(string(data))
	if !is_html {
		t.Errorf("index page should serve html. // %v", string(data))
	}
	if res.StatusCode < http.StatusOK {
		t.Errorf("expected %v got %v ", http.StatusOK, res.Status)
	}
}

func TestGetInvalidPath(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// like a response writer
	w := httptest.NewRecorder()
	// make sure index page serves
	server.GetInvalidPath(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	findhtml := regexp.MustCompile(`\s*<html`)
	is_html := findhtml.Match(data)
	// is_html := findhtml.MatchString(string(data))
	if !is_html {
		t.Errorf("index page should serve html. // %v", string(data))
	}
	if res.StatusCode < http.StatusOK {
		t.Errorf("expected %v got %v ", http.StatusOK, res.Status)
	}
}
