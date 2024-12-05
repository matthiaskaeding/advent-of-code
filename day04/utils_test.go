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
	got, _ := mat.GetCol(0)
	want := "AAS"
	if got != want {
		t.Errorf("Col 0 must be %v but is %v", want, got)
	}
	got, _ = mat.GetCol(1)
	want = "BFJ"
	if got != want {
		t.Errorf("Col 0 must be %v but is %v", want, got)
	}
	got = mat.GetDiagonal()
	want = "AFK"
	if got != want {
		t.Errorf("Col 0 must be %v but is %v", want, got)
	}

	// Matrix:
	// ABC
	// DEF
	// GHI
	lines = []string{"ABC", "DEF", "GHI"}
	mat, _ = NewMat(lines)
	got, _ = mat.GetSubDiagonal(0, 0, "rightdown")
	want = "AEI"
	if got != want {
		t.Errorf("Must be %v but is %v", want, got)
	}
	got, _ = mat.GetSubDiagonal(0, 1, "leftdown")
	want = "BD"
	if got != want {
		t.Errorf("Must be %v but is %v", want, got)
	}

	got, _ = mat.GetSubDiagonal(2, 1, "rightup")
	want = "HF"
	if got != want {
		t.Errorf("Must be %v but is %v", want, got)
	}
	got, _ = mat.GetSubDiagonal(2, 2, "rightup")
	want = "I"
	if got != want {
		t.Errorf("Must be %v but is %v", want, got)
	}
	got, _ = mat.GetSubDiagonal(2, 2, "leftup")
	want = "IEA"
	if got != want {
		t.Errorf("Must be %v but is %v", want, got)
	}
	got, _ = mat.GetSubDiagonal(2, 1, "leftup")
	want = "HD"
	if got != want {
		t.Errorf("Must be %v but is %v", want, got)
	}

	// Matrix:
	// ABCD
	// DEFG
	// GHIJ
	lines = []string{"ABCD", "DEFG", "GHIJ"}
	mat, _ = NewMat(lines)

	got, err := mat.GetSubDiagonalShort(0, 0, "rightdown")
	want = "AEI"
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Must be %v but is %v %v", want, got, len(got))
	}
	got, err = mat.GetSubDiagonalShort(0, 1, "rightdown")
	want = "BFJ"
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Must be %v but is %v %v", want, got, len(got))
	}

}
