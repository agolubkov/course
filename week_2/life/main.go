package main

import (
	"fmt"
	"math/rand"
)

const (
	maxM = 10
	maxN = 25
)

type grid [maxM][maxN]int

var field = grid{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
	{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
}

var stage grid

func (g grid) String() string {
	s := ""
	for _, row := range g {
		s += "\xE2\x94\x82"
		for i, cell := range row {
			if i > 0 {
				s += " "
			}
			if cell == 0 {
				s += " "
			} else {
				s += "\xE2\x80\xA2"
			}
		}
		s += "\xE2\x94\x82\n"
	}
	return s
}

func before(x, maxX int) int {
	if x == 0 {
		return maxX - 1
	}
	return x - 1
}

func after(x, maxX int) int {
	if x == maxX-1 {
		return 0
	}
	return x + 1
}

func countNeighbours(m, n int) int {
	return (field[before(m, maxM)][before(n, maxN)] + field[before(m, maxM)][after(n, maxN)] +
		field[after(m, maxM)][before(n, maxN)] + field[after(m, maxM)][after(n, maxN)] +
		field[before(m, maxM)][n] + field[after(m, maxM)][n] +
		field[m][before(n, maxN)] + field[m][after(n, maxN)])

}

func randomize(gr *grid, seedNum int) {
	rand.Seed(int64(seedNum))
	for m, row := range gr {
		for n := range row {
			gr[m][n] = rand.Intn(2)
		}
	}
}

func generateNextStage(stage *grid) {
	for m := 0; m < maxM; m++ {
		for n := 0; n < maxN; n++ {
			neighbours := countNeighbours(m, n)
			switch {
			case field[m][n] == 0 && neighbours == 3:
				stage[m][n] = 1
			case field[m][n] == 1 && (neighbours < 2 || neighbours > 3):
				stage[m][n] = 0
			default:
				stage[m][n] = field[m][n]
			}
		}
	}
}

func main() {
	var seedNum int

	for step := 1; ; step++ {

		if seedNum > 0 {
			fmt.Println("\033[2J\033[H", "Seed=", seedNum, "Step=", step)
		} else {
			fmt.Println("\033[2J\033[H", "Step=", step)
		}
		fmt.Print(field)
		fmt.Print(`  Enter - new step    r - fill randomly (seed=step)
  q - quit            r[Number] - set seed Number > `)

		var input string
		// check for 'quit' command
		if fmt.Scanln(&input); input == "q" {
			break
		}
		// check for 'randomizing the field' command
		if len(input) > 0 && []rune(input)[0] == 'r' {
			if _, err := fmt.Sscanf(input, "r%d", &seedNum); err != nil || seedNum < 1 {
				seedNum = step
			}
			randomize(&stage, seedNum)
			step = 0
		} else {
			generateNextStage(&stage)
		}

		field, stage = stage, field
	}
}
