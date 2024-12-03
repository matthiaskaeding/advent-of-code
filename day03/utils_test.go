package day03

import (
	"testing"
)

func TestFindAllIndicesRegex(t *testing.T) {
	s := "0abc4abcREST"
	got := FindAllIndicesRegex(s, "abc")
	want := []int{1, 5}
	if got[0] != want[0] {
		t.Errorf("Wanted %v got %v", got, want)
	}
	if got[1] != want[1] {
		t.Errorf("Wanted %v got %v", got, want)
	}

	s = "AAA"
	got = FindAllIndicesRegex(s, "abc")
	if len(got) > 0 {
		t.Errorf("Wanted zero len got %v", got)
	}

}
