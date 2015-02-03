package search

import "testing"

func TestBuildKMPTable(t *testing.T) {
	pattern := "ABCDABD"
	text := "ABC ABCDAB ABCDABCDABDE"
	kmp := KMP{}
	kmp.BuildKMPTable(pattern)
	if kmp.Simulate(text, pattern) != 7 {
		t.Error("Answer not correct for KMP")
	}
}
