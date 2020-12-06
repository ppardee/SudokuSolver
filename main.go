package main

import (
	"bufio"
	"fmt"
	"os"
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

	for {
		fmt.Printf("Pass %v", passCounter)
		fmt.Println("----------")

		clearCount := simpleSolve()

		fmt.Printf("First pass cleared %v cells!", clearCount)

		printPuzzle()

		if clearCount == 0 {
			if unknownCount > 0 {
				fmt.Printf("simpleSolve failed with %v cells unsolved. Exiting", unknownCount)
				break
			} else {
				fmt.Println("Solve complete!")
				break
			}
		}
	}
}

func printPuzzle() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nCurrent state of puzzle\n\n")

	for row := 0; row < 9; row++ {
		fmt.Println(puzzle[row])
	}
	fmt.Println("")
	fmt.Println("Press Enter to continue")
	reader.ReadString('\n')
}
func simpleSolve() int {
	startCount := unknownCount
	for row := 1; row < 10; row++ {
		for col := 1; col < 10; col++ {
			if puzzle[row-1][col-1] != 0 {
				fmt.Printf("Cell [%v, %v] started as %v\n", row, col, puzzle[row-1][col-1])
			} else {
				row1, col1, nonant1 := puzzleParser(row, col)

				fmt.Printf("Row1 has %09b\n", row1)
				fmt.Printf("col1 has %09b\n", col1)
				fmt.Printf("nonant1 has %09b\n", nonant1)

				used := row1 | col1 | nonant1
				fmt.Printf("used has %09b\n", used)

				comp := allOnes &^ used
				guesses[row-1][col-1] = comp
				p := bitsToInts(comp)
				if len(p) == 1 {
					fmt.Printf("Setting Cell [%v, %v] to %v\n", row, col, p[0])
					puzzle[row-1][col-1] = p[0]
					unknownCount--
					fmt.Printf("Unknown count now %v\n", unknownCount)
				} else {
					fmt.Printf("Cell [%v, %v] could be %v\n", row, col, bitsToInts(comp))
				}
			}
		}
	}
	return startCount - unknownCount
}
func puzzleParser(row int, column int) (int, int, int) {
	var colBits int
	var rowBits int
	var nonantBits int

	for i := 0; i < 9; i++ {
		colBits = colBits | intToBits(puzzle[i][column-1])
	}
	for i := 0; i < 9; i++ {
		rowBits = rowBits | intToBits(puzzle[row-1][i])
	}

	var nonantRowStart int
	var nonantColStart int

	if row <= 3 {
		// nonant 1
		nonantRowStart = 0
	} else if row > 3 && row <= 6 {
		// nonant 2
		nonantRowStart = 3
	} else {
		// nonant 2
		nonantRowStart = 6
	}

	if column <= 3 {
		// nonant 1
		nonantColStart = 0
	} else if column > 3 && column <= 6 {
		// nonant 2
		nonantColStart = 3
	} else {
		// nonant 2
		nonantColStart = 6
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			nonantBits = nonantBits | intToBits(puzzle[nonantRowStart+i][nonantColStart+j])
		}
	}

	return rowBits, colBits, nonantBits
}

func intToBits(num int) int {
	switch num {
	case 1:
		return one
	case 2:
		return two
	case 3:
		return three
	case 4:
		return four
	case 5:
		return five
	case 6:
		return six
	case 7:
		return seven
	case 8:
		return eight
	case 9:
		return nine
	case 0:
		return 0
	}
	panic("Out of bounds")
}

func bitsToInts(num int) []int {

	var possibilities []int

	if num|one == num {
		possibilities = append(possibilities, 1)
	}
	if num|two == num {
		possibilities = append(possibilities, 2)
	}
	if num|three == num {
		possibilities = append(possibilities, 3)
	}

	if num|four == num {
		possibilities = append(possibilities, 4)
	}

	if num|five == num {
		possibilities = append(possibilities, 5)
	}

	if num|six == num {
		possibilities = append(possibilities, 6)
	}

	if num|seven == num {
		possibilities = append(possibilities, 7)
	}

	if num|eight == num {
		possibilities = append(possibilities, 8)
	}

	if num|nine == num {
		possibilities = append(possibilities, 9)
	}

	return possibilities
}

func init() {

	// This is the initial state of the puzzle we're trying to solve
	puzzle = [9][9]int{
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

	// Initializing guesses with values provided from the puzzle
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			guesses[i][j] = intToBits(puzzle[i][j])

			if puzzle[i][j] == 0 {
				// Counting the known unknowns!
				unknownCount++
			}
		}
	}
}
