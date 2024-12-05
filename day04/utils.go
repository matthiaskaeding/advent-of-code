package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mat struct {
	lines []string
	n     int
	k     int
}

func NewMat(lines []string) (Mat, error) {
	n := len(lines)
	k := len(lines[0])
	for i := 1; i < n; i++ {
		if len(lines[i]) != k {
			return Mat{}, fmt.Errorf("line %v does not have the same width as first line", i)
		}
	}
	return Mat{lines, n, k}, nil
}

func (m Mat) GetCol(i int) (string, error) {
	if i >= m.k || i < 0 {
		return "", fmt.Errorf("i must be in [0, m.k]")
	}
	var col string
	for j := 0; j < m.n; j++ {
		line := m.lines[j]
		el := line[i]
		col += string(el)
	}
	return col, nil
}
func (m Mat) GetRow(i int) (string, error) {
	if i >= m.k || i < 0 {
		return "", fmt.Errorf("i must be in [0, m.k]")
	}

	return m.lines[i], nil
}
func (m Mat) GetDiagonal() string {
	var diag string
	for j := 0; j < m.n; j++ {
		line := m.lines[j]
		el := line[j]
		diag += string(el)
	}
	return diag
}

// Starting from row 0, at coln, get subdiagiona
// So for instance for this matrix
// ABC
// DEF
// GHI
// The SubDiagonal for coln = 1 would be BD
// The SubDiagonal for coln = 2 would be CEG (the diagional)
// i = row start point
// j = column start point
// direction = where to go from
func (m Mat) GetSubDiagonal(i int, j int, direction string) (string, error) {
	if j >= m.k || j < 0 {
		return "", fmt.Errorf("i must be in [0, m.k]")
	}
	var subDiag string
	var di, dj int
	// Define where to go from start point based on direction
	switch direction {
	case "leftdown":
		di = 1  // rowwise down
		dj = -1 // colwise right
	case "rightdown":
		di = 1
		dj = 1
	case "leftup":
		di = -1
		dj = -1
	case "rightup":
		di = -1
		dj = 1
	default:
		return subDiag,
			fmt.Errorf("direction must be in (leftdown, rightdown, leftup, rightup)")
	}

	for k := 0; k < m.n; k++ {
		el := string(m.lines[i][j])
		subDiag += el
		i += di
		j += dj

		if i < 0 || i >= m.n || j < 0 || j >= m.n {
			break
		}

	}
	return subDiag, nil
}

// Starting from row 0, at coln, get subdiagiona
// So for instance for this matrix
// ABC
// DEF
// GHI
// The SubDiagonal for coln = 1 would be BD
// The SubDiagonal for coln = 2 would be CEG (the diagional)
// i = row start point
// j = column start point
// direction = where to go from
func (m Mat) GetSubDiagonalShort(i int, j int, direction string) (string, error) {
	if i >= m.k || i < 0 {
		return "", fmt.Errorf("i must be in [0, m.k] but is %v. m.k is %v", i, m.k)
	}
	if j >= m.k || j < 0 {
		return "", fmt.Errorf("j must be in [0, m.k] but is %v", j)
	}
	var subDiag string
	var di, dj int
	// Define where to go from start point based on direction
	switch direction {
	case "leftdown":
		di = 1  // rowwise down
		dj = -1 // colwise right
	case "rightdown":
		di = 1
		dj = 1
	case "leftup":
		di = -1
		dj = -1
	case "rightup":
		di = -1
		dj = 1
	default:
		return subDiag,
			fmt.Errorf("direction must be in (leftdown, rightdown, leftup, rightup)")
	}
	for i >= 0 && i < m.n && j >= 0 && j < m.k && len(subDiag) < 3 {
		el := string(m.lines[i][j])
		subDiag += el
		i += di
		j += dj
	}
	return subDiag, nil
}

func ReadLines(file string) (Mat, error) {

	var lines []string
	var mat Mat
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		return mat, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line_i := scanner.Text()
		lines = append(lines, line_i)

	}

	return NewMat(lines)

}

func CountXmas(s string) int {
	return strings.Count(s, "XMAS") + strings.Count(s, "SAMX")
}
func GotMas(s string) bool {
	return (strings.Count(s, "MAS") + strings.Count(s, "SAM")) > 0
}

func Solve() {
	mat, err := ReadLines("day04/input.txt")
	if err != nil {
		panic(err)
	}

	// Simply shift cell by cell and count
	num_xmas := 0
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.k; j++ {
			s0, _ := mat.GetSubDiagonalShort(i, j, "rightdown")
			s1, _ := mat.GetSubDiagonalShort(i, j+2, "leftdown")
			if GotMas(s0) && GotMas(s1) {
				num_xmas++
			}
		}
	}

	fmt.Printf("Solution is: %v\n", num_xmas)
}
