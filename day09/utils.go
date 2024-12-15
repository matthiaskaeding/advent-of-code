package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
func (d Disc) Index(i int) int {
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
func (d Disc) FindSpace() (int, int, bool) {
	u := -1
	for i := range d.data {
		for j := i; j < d.Len() && d.Index(j) == -1; j++ {
			u = j
		}
		if u != -1 {
			return i, u, true
		}
	}
	return 0, 0, false
}

func IsNumber(r byte) bool {
	return r >= '0' && r <= '9'
}

// Returns val, l, u
func (d Disc) GetReplacement(n int) (int, int, int) {
	var bSize, replacement, l, u int
	var exists bool
	for i := u; i >= 0; i-- {
		b := d.data[i]
		if b == -1 {
			continue
		}
		bSize, exists = d.sizes[b]
		if !exists {
			continue
		}
		if bSize > n {
			delete(d.sizes, b)
			continue
		}
		u = i
		l = u - bSize + 1
		replacement = b
	}
	return replacement, l, u
}

func (d Disc) Move() (Disc, bool) {
	l, u, found := d.FindSpace()
	var repl, lRepl, uRepl int
	if found {
		repl, lRepl, uRepl = d.GetReplacement(u - l + 1)
	}

	// Create a new slice with its own underlying array
	db := make([]int, len(d.data))
	copy(db, d.data) // Copy the elements from d to db

	sizeFound := uRepl - lRepl + 1
	for i := 0; i <= sizeFound; i++ {
		posReplacement := lRepl + i
		if posReplacement < l {
			return Disc{db, d.sizes}, true
		}
		db[l+i] = repl
		db[posReplacement] = -1
	}
	delete(d.sizes, repl)

	return Disc{db, d.sizes}, false
}

func (d Disc) MoveAll() Disc {
	d, done := d.Move()
	for !done {
		d, done = d.Move()
	}
	return d
}

func (d Disc) CompCheckSum() int {
	checkSum := 0
	for i, x := range d.data {
		if x == -1 {
			break
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
	discMoved := disc.MoveAll()
	discMoved.Show()
	checkSum := discMoved.CompCheckSum()
	fmt.Printf("Checksum: %v\n", checkSum)

}
