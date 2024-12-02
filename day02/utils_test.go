package day02

import (
	"fmt"
	"testing"
)

func TestLevelCheck(t *testing.T) {
	var l Report
	for i := 0; i < 10; i++ {
		l = append(l, i)
	}
	got := l.Check()
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	l = Report{1, 2}
	got = l.Check()
	if !got {
		t.Errorf("got %v, wanted true", got)
	}

	l = Report{1, 2, 1}
	if got = l.Check(); got {
		t.Errorf("got %v, wanted false, %v", got, l)
	}
	l = Report{2, 1}
	if got = l.Check(); !got {
		t.Errorf("got %v, wanted true, Report: %v", got, l)
	}

	l = Report{2, 1, 5}
	if got = l.Check(); got {
		t.Errorf("got %v, wanted false, Report: %v", got, l)
	}

	l = Report{10, 1}
	if got = l.Check(); got {
		t.Errorf("got %v, wanted false, level: %v", got, l)
	}

}

func testEq(a, b Report) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestLevelDelete(t *testing.T) {
	l := Report{1, 2, 3}
	fmt.Printf("l before removal: %v\n", l)

	got := l.RemoveLevel(0)
	fmt.Printf("l after removal: %v\n", l)

	want := Report{2, 3}

	if !testEq(got, want) {
		t.Errorf("got %v, wanted %v", got, want)

	}

	fmt.Printf("l after removal: %v", l)
	l = Report{1, 2, 3} // Reset l to original

	got = l.RemoveLevel(1)
	want = Report{1, 3}

	if !testEq(got, want) {
		t.Errorf("got %v, wanted %v", got, want)

	}

}

func TestCheckWithRemovel(t *testing.T) {
	l := Report{1, 2, 3}
	got := l.CheckWithRemoval()
	if !got {
		t.Errorf("got %v, wanted true, report: %v", got, l)
	}

	l = Report{1, 10, 3}
	got = l.CheckWithRemoval()
	if !got {
		t.Errorf("got %v, wanted true, report: %v", got, l)
	}

	l = Report{1, 10, 11}
	got = l.CheckWithRemoval()
	if !got {
		t.Errorf("Wanted true, report: %v", l)
	}

	l = Report{999, 10, 11}
	got = l.CheckWithRemoval()
	if !got {
		t.Errorf("Wanted true, report: %v", l)
	}

	l = Report{999, 10, 11, 849384934}
	got = l.CheckWithRemoval()
	if got {
		t.Errorf("Wanted false, report: %v", l)
	}

}
