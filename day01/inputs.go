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
