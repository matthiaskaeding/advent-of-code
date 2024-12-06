package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	first  int
	second int
}

type PairMap map[Pair]bool

// Page ordering: for each number n maps the number that must be before n
type PageOrdering struct {
	pagesBefore map[int][]int
	pagesAfter  map[int][]int
}

func (p PageOrdering) GetPagesBefore(k int) ([]int, bool) {
	o, exists := p.pagesBefore[k]
	return o, exists
}

func (p PageOrdering) GetPagesAfter(k int) ([]int, bool) {
	o, exists := p.pagesAfter[k]
	return o, exists
}

type Update []int

type Updates struct {
	data          []Update
	pagesOrdering PageOrdering
	pairMap       PairMap
}

func (u Updates) GetSortedUpdate(i int) Update {
	ori := u.GetUpdate(i)
	cpy := make(Update, len(ori))
	copy(cpy, ori)
	pairMap := u.pairMap

	sort.Slice(cpy, func(i, j int) bool {
		pageI, pageJ := cpy[i], cpy[j]
		pair := Pair{pageI, pageJ}
		pairIsSorted := pairMap[pair]
		if pairIsSorted {
			return true
		}
		reversedPair := Pair{pageJ, pageI}
		reversedPairIsSorted := pairMap[reversedPair]
		if reversedPairIsSorted {
			return false
		}
		return false
	})

	return cpy
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

func (u Update) isValidUpdate(pageOrdering PageOrdering) bool {
	for _, page := range u {
		pagesBefore, exists := pageOrdering.GetPagesBefore(page)
		if !exists {
			continue
		}
		indAfter := slices.Index(u, page)
		for _, pageBefore := range pagesBefore {
			indBefore := slices.Index(u, pageBefore) // -1 if not present so its fine
			if indBefore > indAfter {
				return false
			}
		}

	}
	return true
}

func (updates Updates) IsValidUpdate(i int) bool {
	u := updates.GetUpdate(i)
	return u.isValidUpdate(updates.pagesOrdering)
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

func (updates Updates) GetSumInValidUpdates() int {
	sum := 0
	for i := 0; i < updates.Len(); i++ {
		isValid := updates.IsValidUpdate(i)
		if !isValid {
			middleVal := updates.GetSortedUpdate(i).GetMiddleVal()
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

	pagesBefore := make(map[int][]int)
	pagesAfter := make(map[int][]int)
	pairMap := make(PairMap)

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
			first, err := strconv.Atoi(parts[0])
			if err != nil {
				return updates, err
			}
			second, err := strconv.Atoi(parts[1])
			if err != nil {
				return updates, err
			}
			pagesBefore[second] = append(pagesBefore[second], first)
			pagesAfter[first] = append(pagesAfter[first], second)

			pair := Pair{first, second}
			//fmt.Printf("Pair: %v, line: %v\n", pair, line)
			pairMap[pair] = true

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
	pageOrdering := PageOrdering{pagesBefore, pagesAfter}
	updates = Updates{updatesData, pageOrdering, pairMap}
	return updates, nil
}

func Solve() {
	updates, err := ReadInput("day05/input.txt")
	if err != nil {
		panic(err)
	}

	u := updates.GetUpdate(4)
	uSorted := updates.GetSortedUpdate(4)
	fmt.Println(u)
	fmt.Println(uSorted)

	sum := updates.GetSumInValidUpdates()
	fmt.Printf("sum: %v\n", sum)

	// Do something
}
