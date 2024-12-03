package day03

import (
	"testing"
)

func TestFindAllIndicesRegex(t *testing.T) {
	s := "0abc4abcREST"
	got := FindAllIndices(s, "abc", true)
	want := []int{1, 5}
	if got[0] != want[0] {
		t.Errorf("Wanted %v got %v", got, want)
	}
	if got[1] != want[1] {
		t.Errorf("Wanted %v got %v", got, want)
	}

	s = "AAA"
	got = FindAllIndices(s, "abc", true)
	if len(got) > 0 {
		t.Errorf("Wanted zero len got %v", got)
	}

}

func TestGetRollInd(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6}
	got, found := GetRollVal(x, 1)
	want := 1
	if got != want || !found {
		t.Errorf("Wanted (%v, true) got (%v, %v)", want, got, found)
	}

	x = []int{0, 1, 10, 12}
	got, found = GetRollVal(x, 11)
	want = 10
	if got != want || !found {
		t.Errorf("Wanted (%v, true) got (%v, %v)", want, got, found)
	}

	x = []int{0, 1, 10, 12}
	got, found = GetRollVal(x, 20)
	want = 12
	if got != want || !found {
		t.Errorf("Wanted (%v, true) got (%v, %v)", want, got, found)
	}

	x = []int{3, 5}
	got, found = GetRollVal(x, 1)
	want = 0
	if got != want || found {
		t.Errorf("Wanted (%v, true) got (%v, %v)", want, got, found)
	}

}
