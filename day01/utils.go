package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func ReadIds(file string) ([]int, []int, error) {
	// Read lines, loop over them and get the ids
	var ids0 []int
	var ids1 []int
	fmt.Println(file)

	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		return ids0, ids1, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	re := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		line := scanner.Text()
		nums := re.FindAllString(line, -1)

		id0, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		ids0 = append(ids0, id0)

		id1, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		ids1 = append(ids1, id1)

	}

	// Return sorted ids
	slices.Sort(ids0)
	slices.Sort(ids1)

	return ids0, ids1, nil

}

func SumDistances(ids0 []int, ids1 []int, e error) int {

	var sum int

	for i, id0 := range ids0 {
		id1 := ids1[i]
		if id1 > id0 {
			sum += id1 - id0
		} else if id0 > id1 {
			sum += id0 - id1
		}
	}
	return sum

}

func CountElement(ids []int, el int) int {
	// Counts how often el appears in slice ids
	// ids must be ordered
	pos, is_present := slices.BinarySearch(ids, el)
	if !is_present {
		return 0
	}
	var count int = 1
	for i := pos + 1; i < len(ids)-1; i++ {
		if ids[i] == el {
			count += 1
		} else {
			break
		}
	}
	return count

}

func CompSimilarityScore(ids0 []int, ids1 []int) int {
	var score int
	for _, id0 := range ids0 {
		count := CountElement(ids1, id0)
		score += id0 * count
	}
	return score

}
