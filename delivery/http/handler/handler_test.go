package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAbout tests GET /about request handler
func TestAbout(t *testing.T) {

	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/about.html", nil)
	if err != nil {
		t.Fatal(err)
	}

	About(httprr, req)
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "About" {
		t.Errorf("want the body to contain the word %q", "about")
	}
}