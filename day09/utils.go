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

type Disc []int

func MakeDisc(s string) Disc {
	var disc Disc
	var id, right int
	for i, r := range s {
		blockSize := int(r - '0')
		if i%2 == 0 {
			right = id
			id++
		} else {
			right = -1
		}
		for j := 0; j < blockSize; j++ {
			disc = append(disc, right)
		}
	}
	return disc
}

func (d Disc) Show() {
	s := ""
	for _, x := range d {
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
	for i := range d {
		for j := i; j < len(d) && d[j] == -1; j++ {
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

func (d Disc) GetReplacement(n int, u int) ([]int, []int) {
	indices := []int{}
	replacements := []int{}
	for i := u; i >= 0 && len(indices) < n; i-- {
		b := d[i]
		if b >= 0 {
			indices = append(indices, i)
			replacements = append(replacements, b)
		}
	}
	return indices, replacements
}

func (d Disc) Move() (Disc, bool) {
	l, u, found := d.FindSpace()
	var (
		indices []int
	)
	if found {
		indices, _ = d.GetReplacement(u-l+1, len(d)-1)
	}

	// Create a new slice with its own underlying array
	db := make([]int, len(d))
	copy(db, d) // Copy the elements from d to db

	for i := l; i <= u; i++ {
		posReplacement := indices[i-l]
		if posReplacement < l {
			return Disc(db), true
		}
		db[i] = db[posReplacement]
		db[posReplacement] = -1
	}
	return Disc(db), false
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
	for i, x := range d {
		if x == -1 {
			break
		}
		checkSum += x * i
	}
	return checkSum
}

func Solve() {
	line := ReadInput("day09/input.txt")
	disc := MakeDisc(line)

	fmt.Println()
	disc.Show()
	discMoved := disc.MoveAll()
	discMoved.Show()
	checkSum := discMoved.CompCheckSum()
	fmt.Printf("Checksum: %v\n", checkSum)

}
