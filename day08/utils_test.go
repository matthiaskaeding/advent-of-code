//go:build exclude

package dayXX

import (
	"testing"
)

func TestFun(t *testing.T) {
	// Test something
	got := l.Check()
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
