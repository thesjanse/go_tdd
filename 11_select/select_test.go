package selecto

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test delayed server", func(t *testing.T) {
		slowServer := makeServer(2)
		fastServer := makeServer(0)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		actual, _ := Racer(slowURL, fastURL)

		if actual != expected {
			t.Errorf("Actual: '%s'; expected: '%s'.", actual, expected)
		}
	})

	t.Run("test timeout", func(t *testing.T) {
		serverA := makeServer(10 * time.Second)
		serverB := makeServer(11 * time.Second)
		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
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
