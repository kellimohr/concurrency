package concurrency

import "testing"

func TestCountWords(t *testing.T) {
	t.Log("Testing the count of words...")

	if (WordCount("one two three four")) != 4 {
		t.Errorf("Expected count of 4")
	}
}
