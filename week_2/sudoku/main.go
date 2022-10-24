package main

import (
	"fmt"
	"math/rand"
	"os"
)

type grid [9][9]int

type cell struct {
	i, j       int
	candidates []int
}

type puzzle struct {
	board         grid
	unknowns      []cell
	seed          int64
	foundSolution bool
}

func (p *puzzle) String() string {
	s := "" + fmt.Sprintln("Seed=", p.seed)
	for i, row := range p.board {
		for j, cell := range row {
			switch j {
			case 0:
				s += "│ "
			case 3, 6:
				s += " │ "
			default:
				s += " "
			}
			switch cell {
			case 0:
				s += " "
			default:
				s += fmt.Sprint(cell)
			}

		}
		s += " │\n"
		if i == 2 || i == 5 {
			s += "├───────┼───────┼───────┤\n"
		}
	}
	return s
}

func (p *puzzle) checkValue(ci, cj, x int) bool {
	for i := 0; i < 9; i++ {
		if p.board[i][cj] == x {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if p.board[ci][j] == x {
			return false
		}
	}
	cornerI := (ci / 3) * 3
	cornerJ := (cj / 3) * 3
	for i := cornerI; i < cornerI+3; i++ {
		for j := cornerJ; j < cornerJ+3; j++ {
			if p.board[i][j] == x {
				return false
			}
		}
	}

	return true
}

func (p *puzzle) genUnknownCells() {
	p.unknowns = make([]cell, 0, 81)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if p.board[i][j] == 0 {
				p.unknowns = append(p.unknowns, cell{i, j, []int{}})
			}
		}
	}
}

func (p *puzzle) findCandidates() {
	for i, c := range p.unknowns {
		for x := 1; x <= 9; x++ {
			if p.checkValue(c.i, c.j, x) {
				p.unknowns[i].candidates = append(p.unknowns[i].candidates, x)
			}
		}
		if len(p.unknowns[i].candidates) == 1 {
			p.board[c.i][c.j] = p.unknowns[i].candidates[0]
		}
	}
}

func (p *puzzle) printCandidates() {
	fmt.Println("Candidates:")
	for i, c := range p.unknowns {
		fmt.Printf("(%v, %v): %-10s ", c.i+1, c.j+1, fmt.Sprintf("%v", c.candidates))
		if (i+1)%4 == 0 || i == len(p.unknowns)-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (p *puzzle) printBoard() {
	fmt.Print("\033[2J\033[H")
	fmt.Println(p)
}

// Brute-force search until the first solution is found
func (p *puzzle) search(n int) {
	c := p.unknowns[n]
	for _, x := range c.candidates {
		if p.checkValue(c.i, c.j, x) {
			p.board[c.i][c.j] = x
			// printBoard()

			if n == len(p.unknowns)-1 {
				p.foundSolution = true
				p.printBoard()
				say("Solution found after brute-force search")
			} else {
				p.search(n + 1)
			}
		}
		if p.foundSolution {
			return
		}
		p.board[c.i][c.j] = 0
	}
}

// Generates a random board (potentially unsolvable)
func (p *puzzle) genRandomBoard() {
	const maxSteps = 50
	p.seed++
	rand.Seed(p.seed)
	p.foundSolution = false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			p.board[i][j] = 0
		}
	}
	for n := 0; n < maxSteps; n++ {
		i := rand.Intn(9)
		j := rand.Intn(9)
		x := rand.Intn(9) + 1
		if p.checkValue(i, j, x) {
			p.board[i][j] = x
		}
	}
}

func (p *puzzle) isSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if p.board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (p *puzzle) hasAnySingleCandidate() bool {
	for i := range p.unknowns {
		if len(p.unknowns[i].candidates) == 1 {
			return true
		}
	}
	return false
}

func (p *puzzle) validateCandidates() bool {
	for i := range p.unknowns {
		if len(p.unknowns[i].candidates) == 0 {
			return false
		}
	}
	return true
}

func say(mes string) {
	if len(mes) > 0 {
		fmt.Println(mes)
	}
	fmt.Print("  <Enter> - continue, <q> - quit  ")
	var ans string
	fmt.Scanln(&ans)
	if ans == "q" {
		os.Exit(0)
	}
}

func main() {
	p := new(puzzle)

randomizeBoard:
	for {
		p.genRandomBoard()

	removingCandidates:
		for {
			p.printBoard()
			p.genUnknownCells()
			if len(p.unknowns) == 0 {
				say("Solution found after removing simple candidates")
				continue randomizeBoard
			}

			p.findCandidates()
			p.printCandidates()

			if !p.validateCandidates() {
				say("No solution")
				continue randomizeBoard
			}

			if p.hasAnySingleCandidate() {
				say("Continue applying single candidates")
				continue removingCandidates
			}

			say("No single candidates found. Attempting brute force search")
			break
		}

		p.search(0)
		if !p.isSolved() {
			say("No solution")
			continue randomizeBoard
		}
	}
}
