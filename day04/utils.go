package day04

import (
	"bufio"
	"fmt"
	"os"
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

func (m Mat) GetCol(i int) string {
	var col string
	for j := 0; j < m.n; j++ {
		line := m.lines[j]
		el := line[i]
		col += string(el)
	}
	return col
}
func (m Mat) GetRow(i int) string {
	return m.lines[i]
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
	col0 := mat.GetCol(0)
	fmt.Printf("\nCol 0 %v\n", col0)
	row0 := mat.GetRow(0)
	fmt.Printf("\nRow 0 %v\n", row0)

}
