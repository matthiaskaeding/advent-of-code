package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const n_reports = 1000

type Report []int
type Reports [n_reports]Report

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
		var report Report
		for _, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			report = append(report, val)
		}
		reports[i] = report
		i++
	}

	if i != n_reports {
		err := fmt.Errorf("expectring 1000 reports, got %v", i)
		return reports, err
	}

	return reports, nil

}

func (l Report) Check() bool {
	n := len(l)
	is_incr := l[1] > l[0]

	for i := 1; i < n; i++ {
		diff := l[i] - l[i-1]
		if is_incr && diff < 0 || !is_incr && diff > 0 {
			return false
		}
		abs_diff := diff
		if !is_incr {
			abs_diff = -diff
		}
		if abs_diff == 0 || abs_diff > 3 {
			return false
		}

	}
	return true

}

func (r Report) RemoveLevel(i int) Report {
	result := make(Report, 0, len(r)-1)
	result = append(result, r[:i]...)
	result = append(result, r[i+1:]...)
	return result
}

func (reports Reports) CountFailures() int {
	var n_failures int
	for _, r := range reports {
		if !r.Check() {
			n_failures++
		}
	}
	return n_failures
}

func Solve() {
	reports, err := ReadReports("day02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(reports[:3])
	n_failures := reports.CountFailures()
	n_safe := n_reports - n_failures
	fmt.Printf("Number of failures: %v, number of safe reports: %v\n", n_failures, n_safe)

}
