package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type interval struct {
	lower int
	upper int
}

func ReadInput(file string) string {

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

	return lines[0]

}

type Disc struct {
	data  []int
	sizes map[int]int // Maps file id to file size
}

func (d Disc) Len() int {
	return len(d.data)
}
func (d Disc) At(i int) int {
	return d.data[i]
}

func MakeDisc(s string) Disc {

	var data []int
	sizes := make(map[int]int)
	var id, right int
	for i, r := range s {
		blockSize := int(r - '0')
		if i%2 == 0 {
			right = id
			sizes[id] = blockSize
			id++
		} else {
			right = -1
		}
		for j := 0; j < blockSize; j++ {
			data = append(data, right)
		}
	}
	return Disc{data, sizes}
}

func (d Disc) Show() {
	s := ""
	for _, x := range d.data {
		if x == -1 {
			s += "."
		} else {
			xs := strconv.Itoa(x)
			s += xs
		}
	}
	fmt.Println(s)
}

// Gives the indices [low, up] for first space
func (d Disc) FindSpace(size int) (interval, bool) {
	for l := range d.data {
		for j := l; j < d.Len() && d.At(j) == -1; j++ {
			u := j
			if u-l+1 == size {
				return interval{l, u}, true
			}
		}
	}
	return interval{}, false
}

func IsNumber(r byte) bool {
	return r >= '0' && r <= '9'
}

// Gets the replacement. For all files where we dont find space,
// we delete the entry in d.sizes, so that we will skip those files in the future
// Returns: replacement value, lower and upper position of space
func (d *Disc) GetReplacement() (interval, interval, int, bool) {
	for i := d.Len() - 1; i >= 0; i-- {
		b := d.At(i)
		if b == -1 {
			continue
		}
		bSize, exists := d.sizes[b]
		if !exists {
			continue
		}
		intervalSpace, foundSpace := d.FindSpace(bSize)
		if !foundSpace || intervalSpace.lower > i {
			delete(d.sizes, b)
			continue
		}
		// i = 10, bsize = 1: i + 1 - 1 = 10 correct
		// i = 10, bsize = 2: i - 2 + 1  = 9  correct
		return intervalSpace, interval{i - bSize + 1, i}, b, true
	}
	return interval{}, interval{}, 0, false
}

func (d *Disc) Move() bool {
	intervalSpace, intervalRepl, repl, replFound := d.GetReplacement()
	if !replFound {
		return false
	}

	intervalSize := intervalSpace.upper - intervalSpace.lower + 1
	for i := 0; i < intervalSize; i++ {
		d.data[intervalSpace.lower+i] = repl
		d.data[intervalRepl.lower+i] = -1

	}

	return true
}

func (d Disc) MoveAll() Disc {
	done := d.Move()
	for done {
		done = d.Move()
	}
	return d
}

func (d Disc) CompCheckSum() int {
	checkSum := 0
	for i, x := range d.data {
		if x == -1 {
			continue
		}
		checkSum += x * i
	}
	return checkSum
}

func Solve() {
	line := ReadInput("day09/input_example.txt")
	disc := MakeDisc(line)

	fmt.Println()
	disc.Show()
	// disc.Move()
	// disc.Show()
	// disc.Move()
	// disc.Show()
	// disc.Move()
	// disc.Show()
	// disc.Move()
	// disc.Show()
	// disc.Move()

	disc.MoveAll()
	disc.Show()
	checkSum := disc.CompCheckSum()
	fmt.Printf("Checksum: %v\n", checkSum)

}
