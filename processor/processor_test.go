package processor

import (
	"os"
	"strings"
	"testing"

	"github.com/kavu/word_counter/counters"
)

func TestProcessor(t *testing.T) {
	cntr := counters.NewStringCounter("123")
	source := strings.NewReader("123asd\n\n123\nzzsd\123x123\n666\n\n\n444\n12")
	expected := 3

	processor := NewProcessor(source, cntr)

	processor.Start()
	processor.Wait()

	if got := processor.GetResult(); got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
	}
}

func TestNewProcessor(t *testing.T) {
	counter := counters.NewStringCounter("123")
	source := os.Stdin

	processor := NewProcessor(source, counter)

	if got := processor.Counter.(*counters.StringCounter); got != counter {
		t.Errorf("Expected %#v, but got %#v", counter, got)
	}

	if got := processor.Source; got != os.Stdin {
		t.Errorf("Expected %#v, but got %#v", os.Stdin, got)
	}
}
