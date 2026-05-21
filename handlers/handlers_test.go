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

// TestHomePageHeroContent verifies the home page includes the design-system hero section.
func TestHomePageHeroContent(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleHome(rec, req); err != nil {
		t.Fatalf("HandleHome returned unexpected error: %v", err)
	}

	body := rec.Body.String()
	wantTexts := []string{
		"Cherished Memories",
		"Add Photo",
		"Explore Stories",
		"Saudade",
	}
	for _, want := range wantTexts {
		if !strings.Contains(body, want) {
			t.Errorf("home page: expected body to contain %q", want)
		}
	}
}

// TestHomePageColorPalette verifies the color palette section names are present.
func TestHomePageColorPalette(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleHome(rec, req); err != nil {
		t.Fatalf("HandleHome returned unexpected error: %v", err)
	}

	body := rec.Body.String()
	palette := []string{"Parchment", "Aged Ink", "Faded Terracotta", "Faded Sage"}
	for _, color := range palette {
		if !strings.Contains(body, color) {
			t.Errorf("home page: expected color palette to contain %q", color)
		}
	}
}

// TestHomePageUIElements verifies the UI elements section is present.
func TestHomePageUIElements(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleHome(rec, req); err != nil {
		t.Fatalf("HandleHome returned unexpected error: %v", err)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "UI Elements") {
		t.Error("home page: expected body to contain \"UI Elements\" section heading")
	}
}

// TestHomePageTitle verifies the HTML title tag is set to "Saudade".
func TestHomePageTitle(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleHome(rec, req); err != nil {
		t.Fatalf("HandleHome returned unexpected error: %v", err)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "<title>Saudade</title>") {
		t.Errorf("home page: expected <title>Saudade</title>, body snippet: %q", body[:min(200, len(body))])
	}
}

// TestLoginPageContent verifies the login page renders the branded form.
func TestLoginPageContent(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleLoginIndex(rec, req); err != nil {
		t.Fatalf("HandleLoginIndex returned unexpected error: %v", err)
	}

	body := rec.Body.String()
	wantTexts := []string{
		"Welcome Back",
		"Sign In",
		`type="email"`,
		`type="password"`,
		`type="submit"`,
		"Saudade",
	}
	for _, want := range wantTexts {
		if !strings.Contains(body, want) {
			t.Errorf("login page: expected body to contain %q", want)
		}
	}
}

// TestLoginPageHasForm verifies the login page includes a POST form to /login.
func TestLoginPageHasForm(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleLoginIndex(rec, req); err != nil {
		t.Fatalf("HandleLoginIndex returned unexpected error: %v", err)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `method="POST"`) {
		t.Error("login page: expected form with method=\"POST\"")
	}
	if !strings.Contains(body, `action="/login"`) {
		t.Error("login page: expected form with action=\"/login\"")
	}
}

// TestNavigationContainsBrand verifies the navigation bar renders the Saudade brand name.
func TestNavigationContainsBrand(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleHome(rec, req); err != nil {
		t.Fatalf("HandleHome returned unexpected error: %v", err)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "Saudade") {
		t.Error("navigation: expected brand name \"Saudade\" in page output")
	}
	if !strings.Contains(body, `href="/login"`) {
		t.Error("navigation: expected login link with href=\"/login\"")
	}
}

// TestNavigationHasHamburgerMenu verifies the navigation includes a hamburger-menu button.
func TestNavigationHasHamburgerMenu(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	if err := handlers.HandleHome(rec, req); err != nil {
		t.Fatalf("HandleHome returned unexpected error: %v", err)
	}

	if !strings.Contains(rec.Body.String(), "Open menu") {
		t.Error("navigation: expected hamburger menu button with aria-label \"Open menu\"")
	}
}

// min is a helper to avoid index out of range in test error messages.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
