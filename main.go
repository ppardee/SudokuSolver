package main

import (
	"fmt"

	"github.com/ppardee/sudokusolver/solver"
)

const one = 1
const two = 1 << 1
const three = 1 << 2
const four = 1 << 3
const five = 1 << 4
const six = 1 << 5
const seven = 1 << 6
const eight = 1 << 7
const nine = 1 << 8

// We're dealing with 9 bits, 255 + 256 = 511
const allOnes = 511

var puzzle solver.SudokuPuzzle

// Creating a guesses matrix for complex solves
var guesses [9][9]int

func main() {
	// We need to know how many cells are left to solve so we know when we are done/stuck
	var unknownCount int

	passCounter := 1
	var ssClearCount int = 0
	var csClearCount int = 0
	for {
		for {
			fmt.Printf("Pass %v\n", passCounter)
			fmt.Println("----------")

			ssClearCount = puzzle.SimpleSolve()

			fmt.Printf("Pass %v cleared %v cells!", passCounter, ssClearCount)

			puzzle.PrintPuzzle()

			if ssClearCount == 0 {
				if puzzle.UnknownCount > 0 {
					fmt.Printf("simpleSolve failed with %v cells unsolved.\n", unknownCount)
					break
				} else {
					fmt.Println("Solve complete!")
					return
				}
			}
			passCounter++
		}

		for {
			fmt.Println("Running complexSolve")
			csClearCount = puzzle.ComplexSolve()
			puzzle.PrintPuzzle()
			if csClearCount == 0 {
				if unknownCount > 0 {
					fmt.Printf("complexSolve failed with %v cells unsolved. Running SimpleSolve again\n", unknownCount)
					ssClearCount = puzzle.SimpleSolve()
					break
				} else {
					fmt.Println("Solve complete!")
					return
				}
			}
		}
		if csClearCount == 0 && ssClearCount == 0 {
			puzzle.PrintPuzzle()
			fmt.Println("Made no progress this round, exiting")
			break
		} else {
			puzzle.PrintPuzzle()
			fmt.Println("Making another pass!")
		}
	}

}

func init() {
	diff := 1
	var p *[9][9]int

	switch diff {
	case 1:
		easy := [9][9]int{
			{0, 8, 0, 7, 0, 9, 0, 0, 2},
			{0, 3, 4, 0, 1, 0, 0, 9, 0},
			{0, 0, 0, 3, 0, 8, 0, 0, 0},
			{0, 0, 6, 4, 3, 0, 8, 0, 1},
			{0, 0, 1, 2, 7, 6, 0, 4, 0},
			{0, 0, 3, 0, 0, 1, 2, 5, 6},
			{0, 0, 0, 0, 9, 0, 0, 2, 7},
			{3, 4, 0, 8, 6, 7, 9, 0, 0},
			{0, 9, 0, 5, 0, 4, 0, 0, 3},
		}
		p = &easy
	case 2:
		medium := [9][9]int{
			{1, 3, 0, 6, 8, 5, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 2},
			{0, 6, 0, 0, 1, 9, 0, 3, 8},
			{0, 0, 1, 0, 0, 0, 0, 4, 0},
			{0, 5, 0, 4, 0, 3, 0, 0, 0},
			{3, 0, 0, 8, 0, 0, 0, 0, 6},
			{4, 2, 7, 5, 6, 0, 9, 0, 0},
			{0, 0, 5, 0, 0, 2, 0, 8, 0},
			{0, 8, 0, 0, 0, 7, 0, 0, 0},
		}
		p = &medium
	default:
		hard := [9][9]int{
			{0, 8, 0, 7, 0, 9, 0, 0, 2},
			{0, 3, 4, 0, 1, 0, 0, 9, 0},
			{0, 0, 0, 3, 0, 8, 0, 0, 0},
			{0, 0, 6, 0, 3, 0, 8, 0, 1},
			{0, 0, 1, 2, 7, 0, 0, 4, 0},
			{0, 0, 3, 0, 0, 1, 2, 5, 6},
			{0, 0, 0, 0, 0, 0, 0, 2, 7},
			{3, 4, 0, 8, 6, 7, 9, 0, 0},
			{0, 9, 0, 5, 0, 4, 0, 0, 3},
		}
		p = &hard
	}

	puzzle = *solver.NewSudokuPuzzle(p)

}
