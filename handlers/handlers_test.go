package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gothstarter/handlers"
)

// TestMake_Success verifies that a handler returning nil results in 200 OK.
func TestMake_Success(t *testing.T) {
	h := handlers.Make(func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

// TestMake_Error verifies that a handler returning an error results in 500.
func TestMake_Error(t *testing.T) {
	h := handlers.Make(func(w http.ResponseWriter, r *http.Request) error {
		return errors.New("something went wrong")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), http.StatusText(http.StatusInternalServerError)) {
		t.Errorf("expected body to contain %q, got %q", http.StatusText(http.StatusInternalServerError), rec.Body.String())
	}
}

// TestHandleHome verifies the home handler returns 200 with an HTML content type.
func TestHandleHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleHome(rec, req); err != nil {
		t.Fatalf("HandleHome returned unexpected error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "text/html") {
		t.Errorf("expected Content-Type to contain text/html, got %q", ct)
	}
}

// TestHandleLoginIndex verifies the login handler returns 200 with an HTML content type.
func TestHandleLoginIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleLoginIndex(rec, req); err != nil {
		t.Fatalf("HandleLoginIndex returned unexpected error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "text/html") {
		t.Errorf("expected Content-Type to contain text/html, got %q", ct)
	}
}
