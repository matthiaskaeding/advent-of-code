package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PageOrdering struct {
	data map[int]int
	n    int
}

func NewPageOrdering(data map[int]int) PageOrdering {
	return PageOrdering{data, len(data)}
}
func (p PageOrdering) GetPageThatMustFollow(k int) (int, bool) {
	PageThatMustFollow, exists := p.data[k]
	return PageThatMustFollow, exists
}

type Update []int

type Updates struct {
	data          []Update
	pagesOrdering PageOrdering
}

func (u Updates) Len() int {
	return len(u.data)
}

func (p Updates) GetUpdate(i int) Update {
	return p.data[i]
}

func (u Update) GetPagesBefore(i int) Update {
	var pagesBefore Update
	if i == 0 {
		return pagesBefore
	}
	for k := 0; k < i; k++ {
		pagesBefore = append(pagesBefore, u[k])
	}
	return pagesBefore
}
func (u Update) GetPagesAfter(i int) Update {
	var pagesAfter Update

	for k := i + 1; k < len(u); k++ {
		pagesAfter = append(pagesAfter, u[k])
	}
	return pagesAfter
}

func (u Update) IsValidUpdate(pageOrdering PageOrdering) bool {
	for j, page := range u {
		pageThatMustFollow, exist := pageOrdering.GetPageThatMustFollow(page)
		if !exist {
			continue
		}
		pagesBefore := u.GetPagesBefore(j)
		if slices.Contains(pagesBefore, pageThatMustFollow) {
			return false
		}
	}
	return true
}

func (updates Updates) IsValidUpdate(i int) bool {
	u := updates.GetUpdate(i)
	return u.IsValidUpdate(updates.pagesOrdering)
}

func (u Update) GetMiddleVal() int {
	n := len(u)
	var i int = (n+1)/2 - 1
	return u[i]
}

func (updates Updates) GetSumValidUpdates() int {
	sum := 0
	for i := 0; i < updates.Len(); i++ {
		isValid := updates.IsValidUpdate(i)
		if isValid {
			middleVal := updates.GetUpdate(i).GetMiddleVal()
			sum += middleVal
		}
	}
	return sum
}

func ReadInput(file string) (Updates, error) {
	var (
		updates     Updates
		updatesData []Update
	)
	pairData := make(map[int]int)

	inFile, err := os.Open(file)
	if err != nil {
		return updates, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			// page ordering
			parts := strings.SplitN(line, "|", 2)
			k, err := strconv.Atoi(parts[0])
			if err != nil {
				return updates, err
			}
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				return updates, err
			}
			pairData[k] = v
		} else if strings.Contains(line, ",") {
			// Pages for update
			parts := strings.Split(line, ",")
			lineData := make(Update, len(parts))

			for i, p := range parts {
				v, err := strconv.Atoi(p)
				if err != nil {
					panic(err)
				}
				lineData[i] = v
			}
			updatesData = append(updatesData, lineData)
		} else {
			continue
		}
	}
	pageOrdering := NewPageOrdering(pairData)
	updates = Updates{updatesData, pageOrdering}
	return updates, nil
}

func Solve() {
	updates, err := ReadInput("day05/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pairs: %v\n", updates.pagesOrdering)
	fmt.Printf("Updates: %v\n", updates.GetUpdate(0))
	middleSum := updates.GetSumValidUpdates()
	fmt.Printf("Middle sum: %v\n", middleSum)

	// Do something
}
