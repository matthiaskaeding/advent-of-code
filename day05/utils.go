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

func NewOrderedPair(data map[int]int) PageOrdering {
	return PageOrdering{data, len(data)}
}
func (p PageOrdering) GetPageThatMustFollow(k int) (int, bool) {
	PageThatMustFollow, exists := p.data[k]
	return PageThatMustFollow, exists
}

type Update []int

type Updates struct {
	data []Update
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

func ReadInput(file string) (PageOrdering, Updates, error) {
	var (
		pairs       PageOrdering
		pageNumbers Updates
	)
	pairData := make(map[int]int)

	inFile, err := os.Open(file)
	if err != nil {
		return pairs, pageNumbers, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			parts := strings.SplitN(line, "|", 2)
			k, err := strconv.Atoi(parts[0])
			if err != nil {
				return pairs, pageNumbers, err
			}
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				return pairs, pageNumbers, err
			}
			pairData[k] = v
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			lineData := make([]int, len(parts))

			for i, p := range parts {
				v, err := strconv.Atoi(p)
				if err != nil {
					panic(err)
				}
				lineData[i] = v
			}
			pageNumbers.data = append(pageNumbers.data, lineData)
		} else {
			continue
		}
	}
	pairs = NewOrderedPair(pairData)
	return pairs, pageNumbers, nil
}

func Solve() {
	pageOrdering, updates, err := ReadInput("day05/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pairs: %v\n", pageOrdering)
	fmt.Printf("Updates: %v\n", updates.GetUpdate(0))

	// Do something
}
