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

func check_diff(x int, x_lag int, is_incr bool) bool {
	diff := x - x_lag
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

	return true

}

func (r Report) Check() bool {
	n := len(r)
	is_incr := r[1] > r[0]

	for i := 1; i < n; i++ {
		x := r[i]
		x_lag := r[i-1]
		ok := check_diff(x, x_lag, is_incr)
		if !ok {
			return false
		}

	}
	return true

}

// 2, 5, 6: Removel of 2 would help
// 2, 5, 6, 7: Removel of 2 would help
// 2, 3, 8: Removel of 8 would help
// 1, 2, 3, 7, 8, 9, 10: Removel of 2 would help
// if check_diff(x[i], x[i-1]) is bad
// then check_diff(x[i+1], x[i-1]) is bad too. Is that true?
func (r Report) CheckWithRemoval() bool {
	n := len(r)
	if n <= 2 {
		return true
	}

	for i := 0; i < n; i++ {
		r_without_i := r.RemoveLevel(i)
		ok := r_without_i.Check()
		if ok {
			return true
		}
	}
	return false

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

func (reports Reports) CountFailuresWithRemovel() int {
	var n_failures int
	for _, r := range reports {
		if !r.CheckWithRemoval() {
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
	fmt.Println("Without removal:")
	fmt.Printf("Number of failures: %v, number of safe reports: %v\n", n_failures, n_safe)

	n_failures_with_removal := reports.CountFailuresWithRemovel()
	n_safe_with_removal := n_reports - n_failures_with_removal
	fmt.Println("With removal:")

	fmt.Printf("Number of failures with: %v, number of safe reports: %v\n", n_failures_with_removal, n_safe_with_removal)

}
