package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GnvSaikiran/snippetbox/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	// checking the Content-Secure-Policy header on the response
	want := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, want, rs.Header.Get("Content-Security-Policy"))

	// checking the other headers on the response
	want = "origin-when-cross-origin"
	assert.Equal(t, want, rs.Header.Get("Referrer-Policy"))

	want = "nosniff"
	assert.Equal(t, want, rs.Header.Get("X-Content-Type-Options"))

	want = "deny"
	assert.Equal(t, want, rs.Header.Get("X-Frame-Options"))

	want = "0"
	assert.Equal(t, want, rs.Header.Get("X-XSS-Protection"))

	// check that the middleware correctly called the next handler in line
	// and the response status code and body are as expected
	assert.Equal(t, http.StatusOK, rs.StatusCode)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	assert.Equal(t, "OK", string(body))
}
