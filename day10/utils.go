package day10

import (
	"bufio"
	"fmt"
	"os"
)

type Mat struct {
	rows [][]int
	n    int
	k    int
}

func NewMat(lines []string) Mat {
	n := len(lines)
	k := len(lines[0])
	rows := make([][]int, n)
	for i := 0; i < n; i++ {
		line := lines[i]
		if len(lines[i]) != k {
			panic("line %v does not have the same width as first line")
		}
		row := make([]int, k)
		for j := 0; j < k; j++ {
			val := int(line[j] - '0')
			row[j] = val
		}
		rows[i] = row
	}
	return Mat{rows, n, k}
}

func (m *Mat) At(i, j int) int {
	return m.rows[i][j]
}

func (m Mat) Move(i, j int, direction string) (int, int) {
	i2 := i
	j2 := j
	switch direction {
	case "u":
		i2--
	case "r":
		j2++
	case "d":
		i2++
	case "l":
		j2--
	default:
		panic("wrong direction")
	}
	// Check for out of bounds
	if i2 < 0 || i2 >= m.n || j2 < 0 || j2 >= m.k {
		return -1, -1
	}
	return i2, j2

}

func (m *Mat) Show() {
	for i := 0; i < m.n; i++ {
		fmt.Println(m.rows[i])
	}
}

func ReadInput(file string) Mat {

	var lines []string
	inFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return NewMat(lines)
}

func Solve() {
	mat := ReadInput("day10/input_example.txt")
	mat.Show()
	// Do something
}
