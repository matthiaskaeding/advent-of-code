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

type Disc string

func MakeDisc(s string) Disc {
	var disc string
	id := 0
	for i, r := range s {
		blockSize := int(r - '0')
		var right string
		if i%2 == 0 {
			right = strconv.Itoa(id)
			id++
		} else {
			right = "."
		}
		for j := 0; j < blockSize; j++ {
			disc += right
		}
	}
	return Disc(disc)
}

// Gives the indices [low, up] for first space
func (d Disc) FindSpace() (int, int, bool) {
	u := -1
	for i := range d {
		for j := i; j < len(d) && d[j] == 46; j++ {
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

func (d Disc) GetReplacement(n int, u int) ([]int, []byte) {
	indices := []int{}
	bytes := []byte{}
	for i := u; i >= 0 && len(indices) < n; i-- {
		b := d[i]
		if IsNumber(b) {
			indices = append(indices, i)
			bytes = append(bytes, b)
		}

	}
	return indices, bytes
}

func (d Disc) Move() (Disc, bool) {
	l, u, found := d.FindSpace()
	var (
		indices []int
	)
	if found {
		indices, _ = d.GetReplacement(u-l+1, len(d)-1)
	}
	db := []byte(d)
	for i := l; i <= u; i++ {
		posReplacement := indices[i-l]
		if posReplacement < l {
			return Disc(db), true
		}
		db[i] = db[posReplacement]
		db[posReplacement] = 46
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
	for i, r := range d {
		if IsNumber(byte(r)) {
			bI := int(r - '0')
			checkSum += i * bI
		} else {
			return checkSum
		}
	}
	return checkSum
}

func Solve() {
	line := ReadInput("day09/input_example.txt")
	fmt.Println(line)
	fmt.Println()

	disc := MakeDisc(line)
	fmt.Println(disc)

	l, u, found := disc.FindSpace()
	fmt.Printf("%v %v %v\n", l, u, found)

	//nmbrs := disc.GetRighmostNumbers(10, len(disc)-1)
	//fmt.Println(nmbrs)
	// Do something
	fmt.Println()
	fmt.Println(disc)
	disc = disc.MoveAll()
	fmt.Println(disc)
	checkSum := disc.CompCheckSum()
	fmt.Printf("Checksum: %v\n", checkSum)

}
