package day02

import "testing"

func TestLevel(t *testing.T) {
	var l Level
	for i := 0; i < 10; i++ {
		l = append(l, i)
	}
	got := l.Check()
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	l2 := Level{1, 2}
	got2 := l2.Check()
	if !got2 {
		t.Errorf("got %v, wanted true", got2)
	}

	l3 := Level{1, 10, 1}
	if got3 := l3.Check(); got3 {
		t.Errorf("got %v, wanted false", got3)
	}
	l4 := Level{10, 1}
	if got4 := l4.Check(); !got4 {
		t.Errorf("got %v, wanted true", got4)
	}

	l5 := Level{10, 1, 100}
	if got5 := l5.Check(); got5 {
		t.Errorf("got %v, wanted false", got5)
	}

}
