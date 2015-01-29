package misc

import "testing"

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome("Anna") {
		t.Error("Palindrome function not working")
	}
}
