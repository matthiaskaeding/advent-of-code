package day08

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Row []string
type Loc [2]int
type Positions map[string][]Loc
type Map struct {
	rows      []Row
	n         int
	k         int
	positions Positions // Maps entries to positions

}

func (l Loc) Diff(other Loc) float64 {
	return math.Sqrt(
		math.Pow(float64(l[0]-other[0]), 2) +
			math.Pow(float64(l[1]-other[1]), 2))
}

func (l Loc) Slope(other Loc) float64 {
	return float64(l[1]-other[1]) / float64(l[0]-other[0])
}

func NewMap(rows []Row) Map {
	positions := make(Positions, 3)
	n := len(rows)
	k := len(rows[0])
	for i := 0; i < len(rows); i++ {
		for j := 0; j < k; j++ {
			s := rows[i][j]
			if s != "." && s != "#" {
				if positions[s] == nil {
					positions[s] = []Loc{}
				}
				positions[s] = append(positions[s], Loc{i, j})
			}
		}
	}
	return Map{rows: rows, positions: positions, n: n, k: k}

}

func (mp Map) IsAntinode(i int, j int) bool {
	l := Loc{i, j}
	for _, locations := range mp.positions {
		differences := make([]float64, 0)
		for _, loc := range locations {
			d := l.Diff(loc)
			differences = append(differences, d)
		}
		if len(differences) == 1 {
			return false
		}
		// differences[i] is the difference of l
		// to locations[i]
		for i := 0; i < len(differences); i++ {
			d := differences[i]
			dd := 2 * d
			j := slices.Index(differences, dd)
			if j != -1 {
				location := locations[i]
				otherLocation := locations[j]
				slope := location.Slope(otherLocation)
				testSlope := location.Slope(l)
				if slope == testSlope {
					return true
				}
			}
		}
	}
	return false
}

func (mp Map) GetAntinodes() []Loc {
	antinodes := make([]Loc, 0)
	for i := 0; i < mp.n; i++ {
		for j := 0; j < mp.n; j++ {
			if mp.IsAntinode(i, j) {
				antinodes = append(antinodes, Loc{i, j})
			}
		}
	}
	return antinodes
}

func ReadInput(file string) ([]string, error) {

	var lines []string
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		return []string{}, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

func parseInput(lines []string) Map {

	rows := make([]Row, len(lines))
	for i := 0; i < len(lines); i++ {
		rw := strings.Split(lines[i], "")
		rows[i] = rw
	}

	return NewMap(rows)
}

func (m Map) Show() {
	for i := 0; i < m.n; i++ {
		fmt.Printf("%v\n", m.rows[i])
	}
}

func Solve() {
	lines, err := ReadInput("day08/input_example.txt")
	if err != nil {
		panic(err)
	}
	mp := parseInput(lines)
	mp.Show()
	fmt.Println()

	fmt.Println()
	tst := mp.IsAntinode(1, 3)
	fmt.Println(tst)

	antinodes := mp.GetAntinodes()
	fmt.Println(antinodes)
	fmt.Printf("Number of antinodes: %v\n", len(antinodes))

}
