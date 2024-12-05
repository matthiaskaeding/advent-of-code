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
	if n != k {
		return Mat{}, fmt.Errorf("input does not give a symmetric matrix")
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

func Solve() {
	mat, err := ReadLines("day04/input.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(mat.lines[i])
	}
	fmt.Printf("\nNum   lines %v", mat.n)
	fmt.Printf("\nWidth lines %v\n", mat.k)
	col0, _ := mat.GetCol(0)
	fmt.Printf("\nCol 0 %v\n", col0)
	row0, _ := mat.GetRow(0)
	fmt.Printf("\nRow 0 %v\n", row0)

	// Now count all the ways
	num_xmas := CountXmas(mat.GetDiagonal())
	for i := 0; i < mat.n; i++ {
		ColI, _ := mat.GetCol(i)
		RowI, _ := mat.GetRow(i)
		num_xmas += CountXmas(ColI)
		num_xmas += CountXmas(RowI)
	}
	for i := 1; i < mat.n; i++ {
		s0, _ := mat.GetSubDiagonal(0, i, "leftdown")
		num_xmas += CountXmas(s0)
		s1, _ := mat.GetSubDiagonal(0, i, "rightdown")
		num_xmas += CountXmas(s1)

		if i != mat.n-1 {
			s2, _ := mat.GetSubDiagonal(mat.n-1, i, "rightup")
			s3, _ := mat.GetSubDiagonal(mat.n-1, i, "leftup")
			num_xmas += CountXmas(s2)
			num_xmas += CountXmas(s3)

		}

	}

	fmt.Printf("Solution is: %v\n", num_xmas)
}
