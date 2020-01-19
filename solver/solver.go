package solver

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	boardSize             = 9
	boardRowSeparator     = "+-------+-------+-------+"
	boardColumnSeparator  = "|"
	boardPlaceholderChar  = "·"
	boardPlaceholderChars = "0_.·"
)

// Board represent a Sudoku board.
type Board [boardSize][boardSize]int

// String display the Board.
func (board Board) String() string {
	builder := strings.Builder{}
	for row := 0; row < boardSize; row++ {
		if row%3 == 0 {
			builder.WriteString(boardRowSeparator + "\n")
		}
		for column := 0; column < boardSize; column++ {
			if column%3 == 0 {
				builder.WriteString(boardColumnSeparator + " ")
			}
			cell := fmt.Sprintf("%d", board[row][column])
			if cell == "0" {
				cell = boardPlaceholderChar
			}
			builder.WriteString(cell + " ")
		}
		builder.WriteString(boardColumnSeparator + "\n")
	}
	builder.WriteString(boardRowSeparator)
	return builder.String()
}

// Solve will return a solution to the puzzle.
func (board *Board) Solve() (*Board, error) {
	solution := board
	solution.backtrack()
	return solution, nil
}

// NewBoard create a new Sudoku board from a string.
func NewBoard(input string) (*Board, error) {
	return NewBoardFrom(strings.NewReader(input))
}

// NewBoardFrom create a new Sudoku board from io.Reader.
func NewBoardFrom(input io.Reader) (*Board, error) {
	board := new(Board)
	row := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == boardRowSeparator {
			continue
		}
		columns := strings.Fields(line)
		if len(columns) == 0 {
			continue
		}
		column := 0
		for _, value := range columns {
			i, err := validCell(value)
			if err == nil {
				board[row][column] = i
				column++
				if column >= boardSize {
					break
				}
			}
		}
		row++
		if row >= boardSize {
			break
		}
	}
	return board, nil
}

// Valid return true if the board is valid.
func (board *Board) Valid() (bool, error) {
	for row := 0; row < boardSize; row++ {
		counter := make(map[int]int)
		for column := 0; column < boardSize; column++ {
			number := board[row][column]
			if number > 0 {
				counter[number]++
				if counter[number] > 1 {
					return false, fmt.Errorf("number %d is duplicated on row %d", number, row+1)
				}
			}
		}
	}

	// Check all columns
	for row := 0; row < boardSize; row++ {
		counter := make(map[int]int)
		for column := 0; column < boardSize; column++ {
			number := board[column][row]
			if number > 0 {
				counter[number]++
				if counter[number] > 1 {
					return false, fmt.Errorf("number %d is duplicated in column %d", number, column+1)
				}
			}
		}
	}

	// Check all regions
	for i := 0; i < boardSize; i += 3 {
		for j := 0; j < boardSize; j += 3 {
			counter := make(map[int]int)
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					number := board[row][col]
					if number > 0 {
						counter[number]++
						if counter[number] > 1 {
							region := fmt.Sprintf("%dx%d", i+1, j+1)
							return false, fmt.Errorf("number %d is duplicated in region %s", number, region)
						}
					}
				}
			}
		}
	}
	return true, nil
}

func validCell(value string) (int, error) {
	if strings.ContainsAny(value, boardPlaceholderChars) {
		return 0, nil
	}
	digit, err := strconv.Atoi(value)
	if err != nil || digit < 1 || digit > boardSize {
		return 0, fmt.Errorf("only digits from 1 to %d and _ as placeholder are allowed values", boardSize)
	}
	return digit, nil
}

func (board *Board) backtrack() bool {
	row, column, solved := board.findEmptyCell()
	if solved {
		return true
	}
	for number := 1; number <= boardSize; number++ {
		if board.isDigitValid(row, column, number) {
			board[row][column] = number
			if board.backtrack() {
				// solution found
				return true
			}
			// failure, reset and try something else
			board[row][column] = 0
		}
	}
	return false
}

func (board *Board) findEmptyCell() (int, int, bool) {
	for row := 0; row < boardSize; row++ {
		for column := 0; column < boardSize; column++ {
			if board[row][column] == 0 {
				return row, column, false
			}
		}
	}
	return 0, 0, true
}

func (board *Board) isDigitValid(row, column, digit int) bool {
	startRow := row / 3 * 3
	startColumn := column / 3 * 3
	for i := 0; i < boardSize; i++ {
		if board[row][i] == digit ||
			board[i][column] == digit ||
			board[startRow+i/3][startColumn+i%3] == digit {
			return false
		}
	}
	return true
}
