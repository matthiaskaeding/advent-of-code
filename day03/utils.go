package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func GetMul(line string) ([]pair, error) {
	sMul := `mul\((\d+),(\d+)\)`

	re := regexp.MustCompile(sMul)
	matches := re.FindAllString(line, -1)

	pairs := make([]pair, len(matches))

	if matches == nil {
		return pairs, fmt.Errorf("no match found in string")
	}

	re_num := regexp.MustCompile(`\d+`)

	for i, m := range matches {
		nums := re_num.FindAllString(m, 2)
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
		pairs[i] = pair
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

func FindAllIndicesRegex(s, substr string) []int {
	var indices []int
	re := regexp.MustCompile(regexp.QuoteMeta(substr))
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
