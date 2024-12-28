package day10

import (
	"testing"
)

func TestComp(t *testing.T) {
	p0 := Point{0, 0}
	p1 := Point{0, 0}
	p2 := Point{0, 1}
	areSame := p0 == p1
	if !areSame {
		t.Errorf("should be the same but are not")
	}
	areSame2 := p0 == p2
	if areSame2 {
		t.Errorf("p0 and p2 must be unequal")
	}

}
