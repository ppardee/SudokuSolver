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

var puzzle [9][9]int

// Creating a guesses matrix for complex solves
var guesses [9][9]int

// We need to know how many cells are left to solve so we know when we are done/stuck
var unknownCount int

func main() {

	passCounter := 1
	var ssClearCount int = 0
	var csClearCount int = 0
	for {
		for {
			fmt.Printf("Pass %v\n", passCounter)
			fmt.Println("----------")

			ssClearCount, unknownCount = solver.SimpleSolve(&puzzle, &guesses)

			fmt.Printf("Pass %v cleared %v cells!", passCounter, ssClearCount)

			solver.PrintPuzzle(&puzzle)

			if ssClearCount == 0 {
				if unknownCount > 0 {
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
			csClearCount, unknownCount = solver.ComplexSolve(&puzzle, &guesses)
			solver.PrintPuzzle(&puzzle)
			if csClearCount == 0 {
				if unknownCount > 0 {
					fmt.Printf("complexSolve failed with %v cells unsolved. Running SimpleSolve again\n", unknownCount)
					ssClearCount, unknownCount = solver.SimpleSolve(&puzzle, &guesses)
					break
				} else {
					fmt.Println("Solve complete!")
					return
				}
			}
		}
		if csClearCount == 0 && ssClearCount == 0 {
			solver.PrintPuzzle(&puzzle)
			fmt.Println("Made no progress this round, exiting")
			break
		} else {
			solver.PrintPuzzle(&puzzle)
			fmt.Println("Making another pass!")
		}
	}

}

func init() {

	// This is the initial state of the puzzle we're trying to solve
	// easy := [9][9]int{
	// 	{0, 8, 0, 7, 0, 9, 0, 0, 2},
	// 	{0, 3, 4, 0, 1, 0, 0, 9, 0},
	// 	{0, 0, 0, 3, 0, 8, 0, 0, 0},
	// 	{0, 0, 6, 4, 3, 0, 8, 0, 1},
	// 	{0, 0, 1, 2, 7, 6, 0, 4, 0},
	// 	{0, 0, 3, 0, 0, 1, 2, 5, 6},
	// 	{0, 0, 0, 0, 9, 0, 0, 2, 7},
	// 	{3, 4, 0, 8, 6, 7, 9, 0, 0},
	// 	{0, 9, 0, 5, 0, 4, 0, 0, 3},
	// }

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

	// hard := [9][9]int{
	// 	{0, 8, 0, 7, 0, 9, 0, 0, 2},
	// 	{0, 3, 4, 0, 1, 0, 0, 9, 0},
	// 	{0, 0, 0, 3, 0, 8, 0, 0, 0},
	// 	{0, 0, 6, 0, 3, 0, 8, 0, 1},
	// 	{0, 0, 1, 2, 7, 0, 0, 4, 0},
	// 	{0, 0, 3, 0, 0, 1, 2, 5, 6},
	// 	{0, 0, 0, 0, 0, 0, 0, 2, 7},
	// 	{3, 4, 0, 8, 6, 7, 9, 0, 0},
	// 	{0, 9, 0, 5, 0, 4, 0, 0, 3},
	// }

	puzzle = medium
	guesses = *solver.InitializeGuesses(&puzzle)

}
