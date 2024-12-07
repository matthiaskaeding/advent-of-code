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

func (d Direction) Change() (Direction, error) {
	switch d {
	case "u":
		return Direction("r"), nil
	case "d":
		return Direction("l"), nil
	case "l":
		return Direction("u"), nil
	case "r":
		return Direction("d"), nil
	default:
		return Direction(""), fmt.Errorf("input directionst be be in u,d,l,r but is %v", d)
	}

}

type GuardMap struct {
	obstacles     ObstacleMatrix
	direction     Direction
	guardPosition GuardPosition
	n             int
	k             int
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
	n := len(obstacles)
	k := len(obstacles[0])
	return GuardMap{obstacles, direction, guardPosition, n, k}, nil
}

func makeDeltaDirection(d Direction) (int, int, error) {
	var di, dj int
	switch d {
	case "u":
		di = -1 // Row up
	case "d":
		di = 1 // Row down
	case "l":
		dj = -1 // Col left
	case "r":
		dj = 1 // Col right
	default:
		return di, dj, fmt.Errorf("wrong value of direction, got %v, must be u, d,l, r", d)
	}
	return di, dj, nil
}

func (guardMap GuardMap) Print(i ...int) {
	var maxLines int
	if i[0] > 0 {
		maxLines = i[0]
	} else {
		maxLines = guardMap.n
	}
	for i := 0; i < maxLines; i++ {
		fmt.Println(guardMap.obstacles[i])
	}
}

func (guardMap *GuardMap) moveGuard() (bool, error) {
	var (
		guardIsFree, hitObstacle bool
	)
	gi := guardMap.guardPosition[0]
	gj := guardMap.guardPosition[1]
	di, dj, err := makeDeltaDirection(guardMap.direction)
	if err != nil {
		return guardIsFree, err
	}

	for !hitObstacle {
		gi = gi + di
		gj = gj + dj
		newPosition := GuardPosition{gi, gj}
		hitObstacle = guardMap.obstacles[gi][gj] == "#"
		if hitObstacle { // Redo moving and change direction
			gi = gi - di
			gj = gj - dj
			newDirection, err := guardMap.direction.Change()
			if err != nil {
				return false, err
			}
			guardMap.direction = newDirection
			guardMap.obstacles[gi][gj] = string(newDirection)
		} else { // Only move when we dont hit an obstacle

			guardMap.obstacles[guardMap.guardPosition[0]][guardMap.guardPosition[1]] = "."
			guardMap.guardPosition = newPosition
		}

		fmt.Printf("gi: %v. gj: %v. Hit obstacle: %v. Direction: %v \n",
			gi, gj, hitObstacle, guardMap.direction)
		// Check if the poor guard is free
		if gi < 0 || gi == guardMap.n || gj < 0 || gj == guardMap.k {
			guardIsFree = true
		}
	}
	return guardIsFree, nil
}

func Solve() {
	guardMap, err := ReadInput("day06/input_example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(guardMap.guardPosition)
	fmt.Println(guardMap.direction)
	fmt.Println(" ")
	fmt.Println(guardMap.direction)
	guardMap.moveGuard()
	fmt.Println(guardMap.direction)

	for i := 0; i < 3; i++ {
		fmt.Println(guardMap.obstacles[i])
	}
	guardMap.moveGuard()
	for i := 0; i < 3; i++ {
		fmt.Println(guardMap.obstacles[i])
	}

	// Do something
}
