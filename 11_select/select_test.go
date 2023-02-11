package selecto

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeServer(2)
	fastServer := makeServer(0)
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	actual := Racer(slowURL, fastURL)

	if actual != expected {
		t.Errorf("Actual: '%s'; expected: '%s'.", actual, expected)
	}
}

func makeServer(delay int) (server *httptest.Server) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delay))
		w.WriteHeader(http.StatusOK)
	}))
	return
}
