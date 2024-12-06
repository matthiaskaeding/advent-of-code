package day05

import (
	"slices"
	"testing"
)

func TestValid(t *testing.T) {
	updates, err := ReadInput("input_example.txt")
	if err != nil {
		t.Error(err)
	}

	u := updates.GetUpdate(0)
	got := u.GetPagesBefore(0)
	if len(got) != 0 {
		t.Errorf("%v must be len 0", got)
	}

	middleValue := updates.GetUpdate(0).GetMiddleVal()
	if middleValue != 61 {
		t.Errorf("Middle value must be 61 but is %v", middleValue)
	}
	middleValue = updates.GetUpdate(2).GetMiddleVal()
	if middleValue != 29 {
		t.Errorf("Middle value must be 29 but is %v", middleValue)
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

	// First 3 are correct, rest is wrong
	for i := 0; i < 6; i++ {
		gotV := updates.IsValidUpdate(i)
		wantV := (i <= 2)
		if gotV != wantV {
			t.Errorf("Update %v. Wanted %v got %v", i, wantV, gotV)
		}
	}

}
