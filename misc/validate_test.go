package misc

import "testing"

func TestValidateEmail(t *testing.T) {
	if !ValidateEmail("me@abhiagarwal.com") {
		t.Error("Email is valid!")
	}
	if ValidateEmail("me@abhiagarwalcom") {
		t.Error("Email is valid!")
	}
}

func TestValidateURL(t *testing.T) {
	if !ValidateURL("http://www.google.com/") {
		t.Error("Email is valid!")
	}
}
