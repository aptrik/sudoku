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

// Solve will return a solution to the puzzle.
func (board *Board) Solve() (Board, error) {
	solution := *board
	solution.backtrack()
	return solution, nil
}

// NewBoard create a new Sudoku board from a string.
func NewBoard(input string) (Board, error) {
	return NewBoardFrom(strings.NewReader(input))
}

// NewBoardFrom create a new Sudoku board from io.Reader.
func NewBoardFrom(input io.Reader) (Board, error) {
	board := Board{}
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
			i, err := board.validCell(value)
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

func (board *Board) validCell(value string) (int, error) {
	digit, err := strconv.Atoi(value)
	if err != nil || digit < 0 || digit > len(board) {
		return 0, fmt.Errorf(
			"only digits from 1 to %d and 0 are allowed values; not %q",
			len(board), value)
	}
	return digit, nil
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
				return true // solution found
			}
			// failure, reset and try something else
			board[row][column] = 0
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

func (board *Board) isPossible(row, column, digit int) bool {
	startRow := row / 3 * 3
	startColumn := column / 3 * 3
	for i := 0; i < len(board); i++ {
		if board[row][i] == digit ||
			board[i][column] == digit ||
			board[startRow+i/3][startColumn+i%3] == digit {
			return false
		}
	}
	return true
}
