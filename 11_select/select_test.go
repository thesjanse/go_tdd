package selecto

import (
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	expected := fastURL
	actual := Racer(slowURL, fastURL)

	if actual != expected {
		t.Errorf("Actual: '%s'; expected: '%s'.", actual, expected)
	}
}
