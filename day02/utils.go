package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Level []int
type Reports []Level

func ReadReports(file string) (Reports, error) {

	var reports Reports

	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		return reports, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	var i int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		if len(nums) < 2 {
			return reports, fmt.Errorf("invalid line format: expected at least 2 numbers, got %d in line '%s'",
				len(nums), line)

		}
		var this_level Level
		for _, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			this_level = append(this_level, val)
		}
		reports = append(reports, this_level)
		i++
	}

	return reports, nil

}

func (l Level) Check() bool {
	n := len(l)
	is_incr := l[1]-l[0] > 0

	for i := 2; i < n; i++ {
		diff := l[i] - l[i-1]
		if is_incr && diff < 0 || !is_incr && diff > 0 {
			return false
		}
	}
	return true

}

func Solve() {
	reports, err := ReadReports("day02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(reports[:3])
}
