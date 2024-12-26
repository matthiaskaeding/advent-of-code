package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var directions = [4]string{"u", "r", "d", "l"}

type Mat struct {
	rows [][]int
	n    int
	k    int
}

type Point struct {
	i, j int
}

func ReadInput(file string) Mat {

	var lines []string
	inFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return NewMat(lines)
}

func NewMat(lines []string) Mat {
	n := len(lines)
	k := len(lines[0])
	rows := make([][]int, n)
	for i := 0; i < n; i++ {
		line := lines[i]
		if len(lines[i]) != k {
			panic("line %v does not have the same width as first line")
		}
		row := make([]int, k)
		for j := 0; j < k; j++ {
			val := int(line[j] - '0')
			row[j] = val
		}
		rows[i] = row
	}
	return Mat{rows, n, k}
}

func (m *Mat) At(i, j int) int {
	return m.rows[i][j]
}

func (m *Mat) IsInBounds(i, j int) bool {
	if i < 0 || i >= m.n || j < 0 || j >= m.k {
		return false
	}
	return true
}

func (m *Mat) GetNeighbors(p Point) []Point {
	neighbors := make([]Point, 0)
	for _, d := range directions {
		p2, inBounds := m.Move(p, d)
		if inBounds {
			neighbors = append(neighbors, p2)
		}
	}
	return neighbors
}

func (m Mat) Move(p Point, direction string) (Point, bool) {
	i2 := p.i
	j2 := p.j
	switch direction {
	case "u":
		i2--
	case "r":
		j2++
	case "d":
		i2++
	case "l":
		j2--
	default:
		panic("wrong direction")
	}
	return Point{i2, j2}, m.IsInBounds(i2, j2)
}

func (m *Mat) Show() {
	for i := 0; i < m.n; i++ {
		fmt.Println(m.rows[i])
	}
}

type Trail map[Point]int

func (m *Mat) checkInner(p Point, dest Point, points Trail) bool {
	val := m.At(p.i, p.j)
	neighbors := m.GetNeighbors(p)
	for _, neighbor := range neighbors {
		valNeigh := m.At(neighbor.i, neighbor.j)
		if neighbor == dest {
			points[neighbor] = len(points)
			fmt.Println("HERE")
			return true
		}
		_, found := points[neighbor]
		if valNeigh == val+1 && !found {
			points[neighbor] = len(points)
			m.checkInner(neighbor, dest, points)
		}
	}
	return false
}

func (m *Mat) Check(p Point, dest Point) (bool, Trail) {
	points := make(Trail, 0)
	canReach := m.checkInner(p, dest, points)

	return canReach, points
}

func (m Mat) ShowTrail(t Trail) {
	for i := 0; i < m.k; i++ {
		row := make([]string, 0)
		for j := 0; j < m.n; j++ {
			p := Point{i, j}
			_, found := t[p]
			if !found {
				row = append(row, ".")
			} else {
				val := m.At(i, j)
				row = append(row, strconv.Itoa(val))
			}
		}
		fmt.Println(row)
	}
}

func (m Mat) GetNines() []Point {
	nines := make([]Point, 0)
	for i := 0; i < m.k; i++ {
		for j := 0; j < m.n; j++ {
			if m.At(i, j) == 9 {
				nines = append(nines, Point{i, j})
			}
		}
	}
	return nines
}

func (m Mat) GetTrailHeadSum() map[Point]int {
	nines := m.GetNines()
	trailHeadSums := make(map[Point]int)
	for i := 0; i < m.k; i++ {
		for j := 0; j < m.n; j++ {
			val := m.At(i, j)
			if val != 0 {
				continue
			}
			src := Point{i, j}
			for _, dest := range nines {
				canReach, _ := m.Check(src, dest)
				if canReach {
					trailHeadSums[src]++
				}
			}
		}
	}
	return trailHeadSums

}

func Solve() {
	mat := ReadInput("day10/input_example.txt")
	mat.Show()

	// canReach, trail := mat.Check(Point{0, 2}, Point{0, 1})
	// fmt.Printf("Can reach destination = %v\nPoints:\n%v\n", canReach, trail)
	// mat.ShowTrail(trail)
	nines := mat.GetNines()
	// fmt.Println(nines)
	// // Do something
	i := 1
	nine := nines[i]
	canReach, trail := mat.Check(Point{0, 4}, nine)
	fmt.Printf("nine[%v] = %v\nCan reach: %v \n", i, nine, canReach)
	mat.ShowTrail(trail)
	trailHeadSum := mat.GetTrailHeadSum()
	fmt.Printf("Trail head sum: %v\n", trailHeadSum)

}
