package day06

import (
	"bufio"
	"fmt"
	"os"
)

// Data model:
// What does it need?
// Direction of guard
// Obstacles
// This can basically be a matrix

type Row []string
type ObstacleRow []string
type ObstacleMatrix []ObstacleRow
type GuardPosition [2]int

type Direction string

func (d Direction) Delta() (int, int, error) {
	switch d {
	case "u":
		return -1, 0, nil
	case "d":
		return 1, 0, nil
	case "l":
		return 0, -1, nil
	case "r":
		return 0, 1, nil
	default:
		return 0, 0, fmt.Errorf("wrong value of direction, got %v, must be u, d,l, r", d)
	}
}

func (d Direction) Change() (Direction, error) {
	switch d {
	case "u":
		return Direction("r"), nil
	case "r":
		return Direction("d"), nil
	case "d":
		return Direction("l"), nil
	case "l":
		return Direction("u"), nil
	default:
		return Direction(""), fmt.Errorf("input direction must be be in u,d,l,r but is %v", d)
	}
}

type Dims struct {
	numRows int
	numCols int
}
type VisitedLocations map[GuardPosition]int

type GuardMap struct {
	obstacles        ObstacleMatrix
	direction        Direction
	guardPosition    GuardPosition
	dims             Dims
	visitedLocations VisitedLocations
}

func (g GuardMap) String() string {
	var s string
	for _, row := range g.obstacles {
		for _, str := range row {
			s += str
		}
		s += "\n"
	}
	return s
}

func ReadInput(file string) (GuardMap, error) {

	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file)
		return GuardMap{}, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	var (
		obstacles     ObstacleMatrix
		direction     Direction
		i             int
		guardPosition GuardPosition
	)

	for scanner.Scan() {
		line := scanner.Text()
		row := make(ObstacleRow, len(line))
		for j, rune := range line {
			switch rune {
			case '^':
				direction = Direction("u")
				row[j] = "u"
				guardPosition = GuardPosition{i, j}
			default:
				row[j] = string(rune)
			}
		}
		obstacles = append(obstacles, row)
		i++
	}
	if direction == "" {
		return GuardMap{}, fmt.Errorf("direction not set")
	}
	dims := Dims{len(obstacles), len(obstacles[0])}
	visitedLocations := make(VisitedLocations)
	visitedLocations[guardPosition] = visitedLocations[guardPosition] + 1
	return GuardMap{obstacles, direction, guardPosition, dims, visitedLocations}, nil
}

func (guardMap *GuardMap) Print(i int) {
	var maxLines int
	if i == -1 {
		maxLines = guardMap.dims.numRows
	} else {
		maxLines = i
	}
	s := "  "
	for i := 0; i < maxLines; i++ {
		fmt.Println(guardMap.obstacles[i])
	}
	fmt.Println(s)
}

func (guardMap *GuardMap) moveGuard() (bool, error) {
	var (
		guardIsFree, hitObstacle bool
	)
	gi := guardMap.guardPosition[0]
	gj := guardMap.guardPosition[1]
	di, dj, err := guardMap.direction.Delta()
	if err != nil {
		return guardIsFree, err
	}

	for !hitObstacle {
		gi = gi + di
		gj = gj + dj
		// Check if the poor guard is free
		if gi < 0 || gi == guardMap.dims.numRows || gj < 0 || gj == guardMap.dims.numCols {
			fmt.Printf("Final position: %v %v\n", gi, gj)
			return true, nil
		}
		hitObstacle = guardMap.obstacles[gi][gj] == "#"
		if hitObstacle { // Redo moving and change direction
			gi = gi - di
			gj = gj - dj
			newDirection, err := guardMap.direction.Change()
			if err != nil {
				return false, err
			}
			guardMap.direction = newDirection
		}
		guardMap.obstacles[gi][gj] = "X"
		guardMap.guardPosition = GuardPosition{gi, gj}
		guardMap.visitedLocations[guardMap.guardPosition] = guardMap.visitedLocations[guardMap.guardPosition] + 1
	}
	return guardIsFree, nil
}

func (guardMap *GuardMap) FreeGuard() (int, error) {
	var (
		numSteps int
		isFree   bool
	)
	for !isFree {
		isFreeI, err := guardMap.moveGuard()
		if err != nil {
			return numSteps, err
		}
		isFree = isFreeI
	}
	return len(guardMap.visitedLocations), nil
}

func Solve() {
	guardMap, err := ReadInput("day06/input_example.txt")
	if err != nil {
		panic(err)
	}

	numSteps, err := guardMap.FreeGuard()
	if err != nil {
		fmt.Print(err)
	}
	guardMap.Print(10)

	fmt.Printf("Number of steps: %v\n", numSteps)

	// Do something
}
