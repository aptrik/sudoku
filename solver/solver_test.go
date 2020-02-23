package solver

import (
	"reflect"
	"testing"
)

var (
	problem1 = Board{
		{0, 0, 5, 0, 0, 0, 4, 0, 9},
		{0, 0, 0, 6, 0, 1, 3, 0, 8},
		{0, 3, 0, 2, 0, 0, 1, 0, 5},
		{5, 0, 4, 0, 1, 0, 6, 0, 0},
		{2, 7, 0, 0, 0, 0, 0, 4, 0},
		{1, 8, 0, 0, 4, 7, 0, 9, 3},
		{7, 0, 0, 0, 3, 9, 0, 1, 6},
		{9, 6, 0, 1, 2, 5, 7, 0, 0},
		{0, 0, 0, 0, 6, 8, 9, 5, 0},
	}
	solution1 = Board{
		{6, 1, 5, 8, 7, 3, 4, 2, 9},
		{4, 2, 9, 6, 5, 1, 3, 7, 8},
		{8, 3, 7, 2, 9, 4, 1, 6, 5},
		{5, 9, 4, 3, 1, 2, 6, 8, 7},
		{2, 7, 3, 9, 8, 6, 5, 4, 1},
		{1, 8, 6, 5, 4, 7, 2, 9, 3},
		{7, 5, 2, 4, 3, 9, 8, 1, 6},
		{9, 6, 8, 1, 2, 5, 7, 3, 4},
		{3, 4, 1, 7, 6, 8, 9, 5, 2},
	}
	problemHardest = Board{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}
	solutionHardest = Board{
		{8, 1, 2, 7, 5, 3, 6, 4, 9},
		{9, 4, 3, 6, 8, 2, 1, 7, 5},
		{6, 7, 5, 4, 9, 1, 2, 8, 3},
		{1, 5, 4, 2, 3, 7, 8, 9, 6},
		{3, 6, 9, 8, 4, 5, 7, 2, 1},
		{2, 8, 7, 1, 6, 9, 5, 3, 4},
		{5, 2, 1, 9, 7, 4, 3, 6, 8},
		{4, 3, 8, 5, 2, 6, 9, 1, 7},
		{7, 9, 6, 3, 1, 8, 4, 5, 2},
	}
)

func TestBoardString(t *testing.T) {
	got := problem1.String()
	want := `0 0 5 0 0 0 4 0 9
0 0 0 6 0 1 3 0 8
0 3 0 2 0 0 1 0 5
5 0 4 0 1 0 6 0 0
2 7 0 0 0 0 0 4 0
1 8 0 0 4 7 0 9 3
7 0 0 0 3 9 0 1 6
9 6 0 1 2 5 7 0 0
0 0 0 0 6 8 9 5 0
`
	if got != want {
		t.Errorf("wrong string representation, got %v, want %v", got, want)
	}
}

func TestBoardNewBoard(t *testing.T) {
	input := problem1.String()
	got, err := NewBoard(input)
	if err != nil {
		t.Errorf("error when creating board %v; got error %v", input, err)
	}
	if !reflect.DeepEqual(got, problem1) {
		t.Errorf("wrong string representation, got %v, want %v", got, problem1)
	}
}

func TestBoardNewBoardInvalid(t *testing.T) {
	input := `
0 x 5 0 0 0 4 0 9
0 0 0 6 0 1 3 0 8
0 3 0 2 0 0 1 0 5
5 0 4 0 1 0 6 0 0
2 7 0 0 0 0 0 4 0
1 8 0 0 4 7 0 9 3
7 0 0 0 3 9 0 1 6
9 6 0 1 2 5 7 0 0
0 0 0 0 6 8 9 5 0
`
	_, err := NewBoard(input)
	if err == nil {
		t.Errorf("board input should not be valid: %v", input)
	}
}

func TestBoardValid(t *testing.T) {
	t.Run("duplicate on row 1", func(t *testing.T) {
		board := Board{
			{1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		if err := board.Valid(); err == nil {
			t.Errorf("expected duplicate number error")
		}
	})
	t.Run("duplicate on row 2", func(t *testing.T) {
		board := Board{
			{0, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		if err := board.Valid(); err == nil {
			t.Errorf("expected duplicate number error")
		}
	})
	t.Run("duplicate in region 3", func(t *testing.T) {
		board := Board{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		if err := board.Valid(); err == nil {
			t.Errorf("expected duplicate number error")
		}
	})
}

func TestBoardSolve(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		got, err := problem1.Solve()
		if err != nil && !reflect.DeepEqual(got, solution1) {
			t.Errorf("solve() got %v, want %v (error %v)", got, solution1, err)
		}
	})
}

func TestBoardSolveInvalid(t *testing.T) {
	t.Run("solve invalid board", func(t *testing.T) {
		puzzle := problem1
		puzzle[0][0] = 5
		_, err := puzzle.Solve()
		if err == nil {
			t.Errorf("invalid board should not be solved: %v", puzzle)
		}
	})
}

func TestSolve_Hardest(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		got, _ := problemHardest.Solve()
		if !reflect.DeepEqual(got, solutionHardest) {
			t.Errorf("solve() got %v, want %v", got, solutionHardest)
		}
	})
}

func BenchmarkSolve(b *testing.B) {
	// BenchmarkSolve-8   	   89318	     12965 ns/op	       0 B/op	       0 allocs/op
	for n := 0; n < b.N; n++ {
		board := problem1
		board.Solve()
	}
}
