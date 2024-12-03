package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type pair struct {
	x    int
	y    int
	prod int
}

func NewPair(x int, y int) pair {
	prod := x * y
	return pair{x, y, prod}
}

func ReadLine(file string) (string, error) {

	var line string

	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		return line, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line_i := scanner.Text()
		line += line_i
	}

	return line, nil

}

// Finds value x <= target minimising target - x
func GetRollVal(x []int, target int) (int, bool) {
	if x[0] > target {
		return 0, false
	}
	i, found := slices.BinarySearch(x, target)
	if found {
		return target, true
	}
	if i == 0 {
		return x[0], true
	}
	return x[i-1], true

}

func GetMul(line string) ([]pair, error) {
	sMul := `mul\((\d+),(\d+)\)`
	reMul := regexp.MustCompile(sMul)

	posMatches := FindAllIndices(line, sMul, false)
	posDont := FindAllIndices(line, "don't()", true)
	posDo := FindAllIndices(line, "do()", true)
	// Because by default do() is activated, we prepend posDo with 0
	posDo = append([]int{0}, posDo...)

	if len(posDont) < 2 {
		panic("There must be at least 2 dont positions")
	}
	if len(posDo) < 2 {
		panic("There must be at least 2 do positions")
	}
	var pairs []pair

	if len(posMatches) == 0 {
		return pairs, fmt.Errorf("no match found in string")
	}

	reNum := regexp.MustCompile(`\d+`)

	for _, posMatch := range posMatches {

		iPosDont, foundDont := GetRollVal(posDont, posMatch)
		iPosDo, foundDo := GetRollVal(posDo, posMatch)

		if iPosDont > iPosDo && foundDont && foundDo {
			continue
		}

		// If we are here its find and we can go on
		mulExpr := reMul.FindString(line[posMatch:])

		nums := reNum.FindAllString(mulExpr, 2)
		if nums == nil {
			return pairs, fmt.Errorf("no numbers")
		}
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		pair := NewPair(x, y)
		pairs = append(pairs, pair)
	}

	return pairs, nil

}

func Prod(pairs []pair) int {
	var prodSum int
	for _, pair := range pairs {
		prodSum += pair.prod
	}
	return prodSum
}

func FindAllIndices(s, substr string, literal bool) []int {
	var indices []int
	var re *regexp.Regexp

	if literal {
		re = regexp.MustCompile(regexp.QuoteMeta(substr))
	} else {
		re = regexp.MustCompile(substr)

	}

	matches := re.FindAllStringIndex(s, -1)

	for _, match := range matches {
		indices = append(indices, match[0])
	}

	return indices
}

func Solve() {
	fmt.Println("DAY 3")
	line, err := ReadLine("day03/input.txt")
	if err != nil {
		panic(err)
	}

	pairs, err := GetMul(line)
	if err != nil {
		panic(err)
	}
	prodSum := Prod(pairs)
	fmt.Println(prodSum)

}
