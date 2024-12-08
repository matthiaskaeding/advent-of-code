package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(file string) ([]string, error) {

	var lines []string
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		return lines, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil

}

type Equation struct {
	lhs []int
	rhs [][]int
}

func ParseLines(lines []string) (Equation, error) {
	lhs := make([]int, len(lines))
	rhs := make([][]int, len(lines))

	for i, line := range lines {
		splitted := strings.SplitN(line, ":", 2)
		lhsI, err := strconv.Atoi(splitted[0])
		if err != nil {
			return Equation{}, fmt.Errorf("failed to parse equation i: %v %v", i, err)
		}
		lhs[i] = lhsI

		rhsStrings := strings.Split(splitted[1], " ")
		for _, rhsString := range rhsStrings {
			rhsString = strings.TrimLeft(rhsString, " ")
			if rhsString == "" {
				continue
			}
			rhsInt, err := strconv.Atoi(rhsString)
			if err != nil {
				return Equation{},
					fmt.Errorf("failed to parse equation i in rhsStrings: %v %v", i, err)
			}
			rhs[i] = append(rhs[i], rhsInt)
		}
	}

	return Equation{lhs, rhs}, nil

}

func Solve() {
	lines, err := ReadInput("day07/input.txt")
	if err != nil {
		panic(err)
	}
	equation, err := ParseLines(lines)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(lines[i])
		fmt.Printf("lhs: %v. rhs: %v\n", equation.lhs[i], equation.rhs[i])
	}

}
