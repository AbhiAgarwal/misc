package search

import "testing"

func compare(X, Y []string) bool {
	for i := 0; i < len(X); i++ {
		if X[i] != Y[i] {
			return false
		}
	}
	return true
}

func TestSuffixTree(t *testing.T) {
	words := []string{
		"aardvark",
		"happy",
		"hello",
		"hero",
		"he",
		"hotel",
	}
	answers := SuffixTree(words, "he")
	matchedWords := []string{
		"hello",
		"hero",
		"he",
	}

	if !compare(answers, matchedWords) {
		t.Error("SuffixTree does not work")
	}
}
