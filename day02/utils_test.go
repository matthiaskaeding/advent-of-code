package day02

import "testing"

func TestLevel(t *testing.T) {
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
