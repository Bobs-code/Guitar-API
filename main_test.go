package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
)

// Tests Home Page Handle
func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	HomePage(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	d := string(b)
	// if err != nil {
	// 	t.Fatalf("expected an integer; got %s", err)
	// }
	if d != "GuitarAPI Project Home Page" {
		t.Fatalf("expected home page welcome; got %v", d)
	}

}
