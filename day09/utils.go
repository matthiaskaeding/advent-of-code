package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInput(file string) Data {

	var lines []string
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if len(lines) != 1 {
		panic("Must be one line")
	}

	return Data(lines[0])

}

type Data string

func (d Data) Get(i int) int {
	s := string(d[i])
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func Solve() {
	data := ReadInput("day09/input_example.txt")
	fmt.Println(data)
	fmt.Println()
	r := data.Get(0)
	fmt.Println(r)

	// Do something
}
