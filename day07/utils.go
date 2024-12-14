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
	lhs int
	rhs []int
}
type Equations []Equation

func ParseLines(lines []string) (Equations, error) {

	var equations Equations

	for i, line := range lines {
		splitted := strings.SplitN(line, ":", 2)
		lhs, err := strconv.Atoi(splitted[0])
		if err != nil {
			return Equations{}, fmt.Errorf("failed to parse equation i: %v %v", i, err)
		}

		rhsStrings := strings.Split(splitted[1], " ")
		rhs := make([]int, 0)
		for _, rhsString := range rhsStrings {
			rhsString = strings.TrimLeft(rhsString, " ")
			if rhsString == "" {
				continue
			}
			rhsInt, err := strconv.Atoi(rhsString)
			if err != nil {
				return Equations{},
					fmt.Errorf("failed to parse equation i in rhsStrings: %v %v", i, err)
			}
			rhs = append(rhs, rhsInt)
		}
		equation := Equation{lhs, rhs}
		equations = append(equations, equation)

	}
	return equations, nil
}

func applyOperator(x int, y int, o string) int {
	switch o {
	case "*":
		return x * y
	case "+":
		return x + y
	case "||":
		out, _ := bangBang(x, y)
		return out
	default:
		return 0
	}
}

func bangBang(x int, y int) (int, error) {
	xs := strconv.Itoa(x)
	ys := strconv.Itoa(y)
	val, err := strconv.Atoi(xs + ys)
	if err != nil {
		return 0, err
	}
	return val, err
}

func genCombinations(lenNumbers int) [][]string {
	// leNnumbers means lenNumbers - 1 operators,
	// which means 2 ^ (lenNumbers - 1) possible combinations
	nOperators := lenNumbers - 1
	nCombinations := 1
	for i := 0; i < nOperators; i++ {
		nCombinations *= 3
	}
	var combinations [][]string

	for i := 0; i < nCombinations; i++ {
		combination := make([]string, nOperators)
		n := i
		for j := 0; j < nOperators; j++ {
			switch n % 3 {
			case 0:
				combination[j] = "*"
			case 1:
				combination[j] = "+"
			case 2:
				combination[j] = "||"
			}
			n /= 3
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

func checkAllCombinationsOfOperators(lhs int, rhs []int) (bool, []string, error) {
	allOperatorCombinations := genCombinations(len(rhs))
	for _, operators := range allOperatorCombinations {
		res, err := applyAllOperators(rhs, operators)
		if err != nil {
			return false, []string{}, err
		}
		if res == lhs {
			return true, operators, nil
		}
	}

	return false, []string{}, nil
}

func Solve() {
	lines, err := ReadInput("day07/input.txt")
	if err != nil {
		panic(err)
	}
	equations, err := ParseLines(lines)
	if err != nil {
		panic(err)
	}

	comb := genCombinations(3)
	fmt.Println(comb)

	sum := 0
	for i := 0; i < len(equations); i++ {
		fmt.Printf("lhs: %v. rhs: %v\n", equations[i].lhs, equations[i].rhs)
		check, operators, err := checkAllCombinationsOfOperators(equations[i].lhs, equations[i].rhs)
		if err != nil {
			panic(err)
		}
		if check {
			sum += equations[i].lhs
		}
		fmt.Printf("Res: %v, %v, \n\n", check, operators)
	}
	fmt.Printf("Res: %v \n", sum)

}
