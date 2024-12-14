package day07

import (
	"testing"
)

func TestFun(t *testing.T) {
	// Test something
	got, _ := bangBang(1, 23)
	want := 123
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
