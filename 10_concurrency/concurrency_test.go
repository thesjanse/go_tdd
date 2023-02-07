package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://dfasdf.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://rambler.ru",
		"waat://dfasdf.geds",
	}

	expected := map[string]bool{
		"http://google.com":  true,
		"http://rambler.ru":  true,
		"waat://dfasdf.geds": false,
	}

	actual := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: '%v'; actual: '%v'.", expected, actual)
	}
}
