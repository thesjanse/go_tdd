package selecto

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test delayed server", func(t *testing.T) {
		slowServer := makeServer(2 * time.Millisecond)
		fastServer := makeServer(1 * time.Microsecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		actual, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Got an unexpected error '%v'", err)
		}

		if actual != expected {
			t.Errorf("Actual: '%s'; expected: '%s'.", actual, expected)
		}
	})

	t.Run("test timeout", func(t *testing.T) {
		serverA := makeServer(10 * time.Millisecond)
		serverB := makeServer(11 * time.Millisecond)
		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 9*time.Millisecond)
		assertError(t, err, ErrorTimeout)
	})
}

func makeServer(delay time.Duration) (server *httptest.Server) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return
}

func assertError(t testing.TB, actual, expected error) {
	t.Helper()
	if actual == nil {
		t.Fatal("Didn't got an error!")
	}
	if actual != expected {
		t.Errorf("Actual: '%q' but expected: '%q", actual, expected)
	}
}
