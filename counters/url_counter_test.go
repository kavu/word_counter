package counters

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLCounter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "123666123")
	}))
	defer server.Close()

	searchString := "123"
	expectedTotal := 2
	counter := NewURLCounter(searchString, 10)

	counter.Count(server.URL)
	counter.Wait()

	if counter.Counted() != expectedTotal {
		t.Errorf("Expected %d, but got %d", expectedTotal, counter.Counted())
	}
}

func TestNewURLCounter(t *testing.T) {
	expected := "123"
	maxJob := 5
	counter := NewURLCounter(expected, maxJob)

	if got := counter.searchString; got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
	}

	if got := cap(counter.maxJobs); got != maxJob {
		t.Errorf("Expected %d, but got %d", expected, got)
	}
}
