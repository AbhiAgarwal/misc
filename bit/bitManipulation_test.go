package bit

import "testing"

func TestIsOn(t *testing.T) {
	if IsOn(01100, 0) == true {
		t.Error("Wrong manipulation")
	}
}

func TestSetBit(t *testing.T) {
	SetBit(01100, 0)
}
