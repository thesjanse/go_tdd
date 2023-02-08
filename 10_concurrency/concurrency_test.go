package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://dfasdf.geds" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
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

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "an url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
