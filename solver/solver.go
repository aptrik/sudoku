package solver

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Board represent a Sudoku board.
type Board [9][9]int

// String display the Board.
func (board Board) String() string {
	builder := strings.Builder{}
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board[row]); column++ {
			cell := fmt.Sprintf("%d", board[row][column])
			builder.WriteString(cell)
			if column < len(board[row])-1 {
				builder.WriteString(" ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

// NewBoard create a new Sudoku board from a string.
func NewBoard(input string) (Board, error) {
	return NewBoardFrom(strings.NewReader(input))
}

// NewBoardFrom create a new Sudoku board from io.Reader.
func NewBoardFrom(input io.Reader) (Board, error) {
	board := Board{}
	validCell := func(value string) (int, error) {
		digit, err := strconv.Atoi(value)
		if err != nil || digit < 0 || digit > len(board) {
			return 0, fmt.Errorf(
				"only digits from 1 to %d and 0 are allowed values; not %q",
				len(board), value)
		}
		return digit, nil
	}

	row := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		columns := strings.Fields(line)
		if len(columns) == 0 {
			continue
		}
		column := 0
		for _, value := range columns {
			i, err := validCell(value)
			if err != nil {
				return board, err
			}
			board[row][column] = i
			column++
			if column >= len(board[row]) {
				break
			}
		}
		row++
		if row >= len(board) {
			break
		}
	}
	return board, nil
}

// Solve will return a solution to the puzzle.
func (board *Board) Solve() (Board, error) {
	puzzle := *board
	if err := puzzle.Valid(); err != nil {
		return puzzle, err
	}
	puzzle.backtrack()
	return puzzle, nil
}

func (board *Board) backtrack() bool {
	row, column := board.findEmptyCell()
	if row < 0 || column < 0 {
		return true
	}
	for number := 1; number <= len(board); number++ {
		if board.isPossible(row, column, number) {
			board[row][column] = number
			if board.backtrack() {
				return true // first solution found
			}
			board[row][column] = 0 // failure, try another number
		}
	}
	return false
}

func (board *Board) findEmptyCell() (int, int) {
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board[row]); column++ {
			if board[row][column] == 0 {
				return row, column
			}
		}
	}
	return -1, -1
}

func (board *Board) isPossible(row, column, number int) bool {
	startRow := row / 3 * 3
	startColumn := column / 3 * 3
	for i := 0; i < len(board); i++ {
		if board[row][i] == number ||
			board[i][column] == number ||
			board[startRow+i/3][startColumn+i%3] == number {
			return false
		}
	}
	return true
}

// Valid checks if the board is a valid puzzle.
func (board *Board) Valid() error {
	// Check all rows
	for row := 0; row < len(board); row++ {
		counter := make(map[int]int)
		for column := 0; column < len(board); column++ {
			number := board[row][column]
			if number > 0 {
				counter[number]++
				if counter[number] > 1 {
					return fmt.Errorf("number %d is duplicated in row %d", number, row+1)
				}
			}
		}
	}

	// Check all columns
	for column := 0; column < len(board); column++ {
		counter := make(map[int]int)
		for row := 0; row < len(board); row++ {
			number := board[row][column]
			if number > 0 {
				counter[number]++
				if counter[number] > 1 {
					return fmt.Errorf("number %d is duplicated in column %d", number, column+1)
				}
			}
		}
	}

	// Check all regions
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			counter := make(map[int]int)
			for row := i; row < i+3; row++ {
				for column := j; column < j+3; column++ {
					number := board[row][column]
					if number > 0 {
						counter[number]++
						if counter[number] > 1 {
							region := fmt.Sprintf("%dx%d", (row+3)/3, (column+3)/3)
							return fmt.Errorf(
								"number %d is duplicated in region %s (row %d, column %d)",
								number, region, row+1, column+1)
						}
					}
				}
			}
		}
	}
	return nil
}
