package solver

import "fmt"

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

// SudokuPuzzle represents the state of a puzzle
type SudokuPuzzle struct {
	Puzzle       [9][9]int
	Guesses      [9][9]int
	UnknownCount int
}

// PrintPuzzle prints the puzzle to the console
func (s *SudokuPuzzle) PrintPuzzle() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nCurrent state of puzzle\n\n")

	for row := 0; row < 9; row++ {
		fmt.Println(s.Puzzle[row])
	}
	fmt.Println("")
	//fmt.Println("Press Enter to continue")
	//reader.ReadString('\n')
}

// SimpleSolve attempts to solve the puzzle and returns the number of unknownCells solved
func (s *SudokuPuzzle) SimpleSolve() int {
	startCount := s.UnknownCount

	for row := 1; row < 10; row++ {
		for col := 1; col < 10; col++ {

			rowIdx, colIdx := row-1, col-1

			if s.Puzzle[rowIdx][colIdx] != 0 {
				// If the value is already set, move on
				//fmt.Printf("Cell [%v, %v] started as %v\n", row, col, puzzle[rowIdx][colIdx])
			} else {
				// Get the bit-wise representation of each row, column and nonant
				rowBits, colBits, nonantBits := puzzleParser(&s.Puzzle, row, col)

				// Bitwise OR together all the results to get a complete list of numbers used in the relevant areas
				used := rowBits | colBits | nonantBits

				// Now that we know what numbers are used, we need to find out which ones aren't used,
				// so we need to flip the 1s to 0s and 0s to 1s.  Doing an Bitwise AND NOT on 111111111 will flip them
				comp := allOnes &^ used

				// Store the numbers that can fit
				s.Guesses[rowIdx][colIdx] = comp

				// Get a slice of the possibilities in numerical representation from the bit representation
				p := bitsToInts(comp)

				if len(p) == 1 {
					// If there is only one possibility, set the cell to that value and decrement the unknown count
					fmt.Printf("Setting Cell [%v, %v] to %v\n", row, col, p[0])
					s.Puzzle[rowIdx][colIdx] = p[0]
					s.UnknownCount--
					fmt.Printf("Unknown count now %v\n", s.UnknownCount)
				} else {
					fmt.Printf("Cell [%v, %v] could be %v\n", row, col, bitsToInts(comp))
				}
			}
		}
	}
	return startCount - s.UnknownCount
}

// ComplexSolve attempts to solve the puzzle and returns the number of unknownCells solved
func (s *SudokuPuzzle) ComplexSolve() int {
	startCount := s.UnknownCount
	// TODO: handle the state where there is a possible but not definitive solution
	// Example:
	// Cell [4, 1] could be [4 9]
	// Cell [4, 4] could be [4 9]
	// Cell [6, 1] could be [4 9]
	// Cell [6, 4] could be [4 9]
	// We could choose one cell to be 9 and the rest would fall in place.

	for row := 1; row < 10; row++ {
		for col := 1; col < 10; col++ {
			rowIdx, colIdx := row-1, col-1

			if s.Puzzle[rowIdx][colIdx] == 0 {
				// First we want to check to see if any of the cell's "siblings" to see if this cell has an exclusive guess

				// Get the bit-wise representation of each row, column and nonant
				rowBits, colBits, nonantBits := 0, 0, 0

				// We need to find all of the unsolved sibling cells and then OR their guess bits together
				for i := 0; i < 9; i++ {
					if i != colIdx && len(bitsToInts(s.Guesses[rowIdx][i])) > 1 {
						rowBits = rowBits | s.Guesses[rowIdx][i]
					}
				}

				for i := 0; i < 9; i++ {
					if i != rowIdx && len(bitsToInts(s.Guesses[i][colIdx])) > 1 {
						colBits = colBits | s.Guesses[i][colIdx]
					}
				}

				nonantRowStart := getNonantStart(row)
				nonantColStart := getNonantStart(col)

				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						nRowIdx := nonantRowStart + i
						nColIdx := nonantColStart + j
						if nRowIdx != rowIdx && nColIdx != colIdx && len(bitsToInts(s.Guesses[nRowIdx][nColIdx])) > 1 {
							nonantBits = nonantBits | s.Guesses[nRowIdx][nColIdx]
						}
					}
				}

				// Bitwise OR together all the results to get a complete list of guesses in the relevant areas
				allGuesses := rowBits | colBits | nonantBits
				thisCellsGuessSlice := bitsToInts(s.Guesses[rowIdx][colIdx])
				var refinedGuesses []int

				for _, v := range thisCellsGuessSlice {
					if !contains(bitsToInts(allGuesses), v) {
						refinedGuesses = append(refinedGuesses, v)
					}
				}

				if len(refinedGuesses) == 1 {
					// If there is only one possibility, set the cell to that value and decrement the unknown count
					fmt.Printf("Setting Cell [%v, %v] to %v\n", row, col, refinedGuesses[0])
					s.Puzzle[rowIdx][colIdx] = refinedGuesses[0]
					s.UnknownCount--
					fmt.Printf("Unknown count now %v\n", s.UnknownCount)
				}
			}
		}
	}
	return startCount - s.UnknownCount
}

// NewSudokuPuzzle returns a populated SudokuPuzzle
func NewSudokuPuzzle(puzzle *[9][9]int) *SudokuPuzzle {
	// Initializing guesses with values provided from the puzzle
	ret := SudokuPuzzle{
		Puzzle:       *puzzle,
		UnknownCount: 0,
	}

	guesses := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			guesses[i][j] = intToBits(puzzle[i][j])

			if puzzle[i][j] == 0 {
				// Counting the known unknowns!
				ret.UnknownCount++
			}
		}
	}
	ret.Guesses = guesses
	return &ret
}

// Gets the bit representation for the row, column and nonant (like a quadrant, but there are nine instead of four segments)
func puzzleParser(puzzle *[9][9]int, row int, column int) (int, int, int) {
	var colBits int
	var rowBits int
	var nonantBits int

	for i := 0; i < 9; i++ {
		colBits = colBits | intToBits(puzzle[i][column-1])
	}

	for i := 0; i < 9; i++ {
		rowBits = rowBits | intToBits(puzzle[row-1][i])
	}

	nonantRowStart := getNonantStart(row)
	nonantColStart := getNonantStart(column)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			nonantBits = nonantBits | intToBits(puzzle[nonantRowStart+i][nonantColStart+j])
		}
	}

	return rowBits, colBits, nonantBits
}

func getNonantStart(index int) int {

	if index <= 3 {
		return 0
	} else if index > 3 && index <= 6 {
		return 3
	}
	return 6

}

// Converts integers to bit representations of the number in an area. Each bit position indicates if that number is present.
// For example, 1 is represented as 000000001, 9 is represented as 100000000 and 4 is represented as 000001000
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

// Reverses the numerical to bit conversion
func bitsToInts(num int) []int {

	var possibilities []int

	if num == 0 {
		return possibilities
	}

	// Bitwise OR will not change a number if the bit representation is present.
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

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
