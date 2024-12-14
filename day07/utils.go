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

func applyOperator(x int, y int, o string) int {
	if o == "*" {
		return x * y
	} else {
		return x + y
	}
}

func genCombinations(lenNumbers int) [][]string {
	// leNnumbers means lenNumbers - 1 operators,
	// which means 2 ^ (lenNumbers - 1) possible combinations
	nOperators := lenNumbers - 1
	nCombinations := 1 << nOperators
	var combinations [][]string

	for i := 0; i < nCombinations; i++ {
		combination := make([]string, nOperators)
		for j := 0; j < nOperators; j++ {
			if (i & (1 << j)) != 0 {
				combination[j] = "*"
			} else {
				combination[j] = "+"
			}
		}
		combinations = append(combinations, combination)
	}
	return combinations
}

// Apply operators left to right
// What does this mean?
// for instane for
// a * b + c
// its not
// a * (b + c)
// but instead (a * b) + c
// a * b + c * d is than
// ((a * b) + c) * d
func applyAllOperators(numbers []int, operators []string) (int, error) {
	if len(numbers)-1 != len(operators) {
		return 0, fmt.Errorf("there must be one less operator then numbers, but are %v, %v",
			len(numbers), len(operators))
	}
	x := applyOperator(numbers[0], numbers[1], operators[0])
	for i := 1; i < len(operators); i++ {
		y := numbers[i+1]
		x = applyOperator(x, y, operators[i])
	}
	return x, nil
}

func 

func Solve() {
	lines, err := ReadInput("day07/input_example.txt")
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

	i := 1
	res, err2 := applyAllOperators(equation.rhs[i], []string{"+", "*"})
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("Lhs: %v, rhs: %v  Res: %v\n", equation.lhs[i], equation.rhs[i], res)

}
