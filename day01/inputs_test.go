package day01

import "testing"

func TestCountElement(t *testing.T) {
	x := []int{1, 2, 3}
	want := 1
	for i := 1; i < 4; i++ {
		got := CountElement(x, i)
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
	got := CountElement(x, 15)
	want = 0
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	x2 := []int{100, 200, 200, 900}
	got2 := CountElement(x2, 200)
	want2 := 2
	if got2 != want2 {
		t.Errorf("got %q, wanted %q", got2, want2)
	}

}
