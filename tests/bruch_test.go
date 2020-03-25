package tests

import (
	"github.com/linalinn/Bruch/bruch"
	"testing"
)

func TestBruch_String(t *testing.T) {
	b1, _ := bruch.NewBruch(90, 20)
	if b1.String() != "4 1/2" {
		t.Fail()
	}

}
