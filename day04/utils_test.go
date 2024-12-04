package day04

import (
	"testing"
)

func TestMat(t *testing.T) {
	lines := []string{"ABC", "AFP", "SJK"}
	mat, _ := NewMat(lines)
	if mat.n != 3 {
		t.Errorf("mat must have 3 rows, has %v", mat.n)
	}
	got := mat.GetCol(0)
	want := "AAS"
	if got != want {
		t.Errorf("Col 0 must be %v but is %v", want, got)
	}
	got = mat.GetCol(1)
	want = "BFJ"
	if got != want {
		t.Errorf("Col 0 must be %v but is %v", want, got)
	}

	lines = []string{"ABC", "AFP", "SJK", "AAR"}
	mat, err := NewMat(lines)
	if err == nil {
		t.Errorf("must throw error but is %v", err)
	}

}
