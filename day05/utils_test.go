package day05

import (
	"slices"
	"testing"
)

func TestValid(t *testing.T) {
	pageOrdering, updates, err := ReadInput("input_example.txt")
	if err != nil {
		t.Error(err)
	}

	u := updates.GetUpdate(0)
	got := u.GetPagesBefore(0)
	if len(got) != 0 {
		t.Errorf("%v must be len 0", got)
	}

	got = u.GetPagesBefore(1)
	want := Update{75}
	if !slices.Equal(got, want) {
		t.Errorf("wanted %v got %v", want, got)
	}
	got = u.GetPagesAfter(0)
	want = Update{47, 61, 53, 29}
	if !slices.Equal(got, want) {
		t.Errorf("wanted %v got %v", want, got)
	}

	gotV := u.IsValidUpdate(pageOrdering)
	wantV := true
	if gotV != wantV {
		t.Errorf("wanted %v got %v", wantV, gotV)
	}

	u = updates.GetUpdate(3)
	gotV = u.IsValidUpdate(pageOrdering)
	wantV = false
	if gotV != wantV {
		t.Errorf("wanted %v got %v", wantV, gotV)
	}

}
