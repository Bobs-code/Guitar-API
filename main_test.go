package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
)

// Tests Home Page Handle
func TestHomePage(t *testing.T) {
	tt := []struct {
		name  string
		value string
		ret   string
		err   string
	}{
		{name: "Home Page Request Test", value: "/", ret: "GuitarAPI Project Home Page"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080"+tc.value, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			homePage(rec, req)
			res := rec.Result()
			defer res.Body.Close()
			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected STatus Bad Request; got %v", res.Status)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			d := string(b)
			// if err != nil {
			// 	t.Fatalf("expected an integer; got %s", err)
			// }
			if d != tc.ret {
				t.Fatalf("expected %v; got %v", tc.ret, d)
			}
		})
	}

}
