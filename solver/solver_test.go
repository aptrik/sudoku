package solver

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	Problem1 = `
0 0 5 0 0 0 4 0 9
0 0 0 6 0 1 3 0 8
0 3 0 2 0 0 1 0 5
5 0 4 0 1 0 6 0 0
2 7 0 0 0 0 0 4 0
1 8 0 0 4 7 0 9 3
7 0 0 0 3 9 0 1 6
9 6 0 1 2 5 7 0 0
0 0 0 0 6 8 9 5 0`
	Problem1Pretty = `+-------+-------+-------+
| · · 5 | · · · | 4 · 9 |
| · · · | 6 · 1 | 3 · 8 |
| · 3 · | 2 · · | 1 · 5 |
+-------+-------+-------+
| 5 · 4 | · 1 · | 6 · · |
| 2 7 · | · · · | · 4 · |
| 1 8 · | · 4 7 | · 9 3 |
+-------+-------+-------+
| 7 · · | · 3 9 | · 1 6 |
| 9 6 · | 1 2 5 | 7 · · |
| · · · | · 6 8 | 9 5 · |
+-------+-------+-------+`
	Problem1Solution = `6 1 5 8 7 3 4 2 9
4 2 9 6 5 1 3 7 8
8 3 7 2 9 4 1 6 5
5 9 4 3 1 2 6 8 7
2 7 3 9 8 6 5 4 1
1 8 6 5 4 7 2 9 3
7 5 2 4 3 9 8 1 6
9 6 8 1 2 5 7 3 4
3 4 1 7 6 8 9 5 2`
)

func Test_NewBoard(t *testing.T) {
	t.Run("pretty", func(t *testing.T) {
		board, _ := NewBoard(Problem1)
		got := board.String()
		if got != Problem1Pretty {
			t.Errorf("NewBoard() got %v, want %v", got, Problem1Pretty)
		}
	})
	t.Run("parse-output", func(t *testing.T) {
		board, _ := NewBoard(Problem1Pretty)
		got := board.String()
		if valid, err := board.Valid(); !valid || err != nil {
			t.Errorf("board is not valid(), got %v, error %v", valid, err)
		}
		if got != Problem1Pretty {
			t.Errorf("NewBoard() got %v, want %v", got, Problem1Pretty)
		}
	})
}

func Test_Solve(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		board, _ := NewBoard(Problem1)
		solution, _ := NewBoard(Problem1Solution)
		if valid, err := solution.Valid(); !valid || err != nil {
			t.Errorf("expected solution board is not valid(), got %v, error %v", valid, err)
		}
		got, _ := board.Solve()
		if valid, err := got.Valid(); !valid || err != nil {
			t.Errorf("solution is not valid(), got %v, error %v", valid, err)
		}
		if !reflect.DeepEqual(got, solution) {
			t.Errorf("solve() got %v, want %v", got, solution)
		}
	})
}

func Test_Valid(t *testing.T) {
	t.Run("duplicate on row 1", func(t *testing.T) {
		board, _ := NewBoard(`
1 1 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0`)
		if valid, err := board.Valid(); valid && err != nil {
			t.Errorf("expected duplicate number error")
		}
	})
	t.Run("duplicate on row 2", func(t *testing.T) {
		board, _ := NewBoard(`
0 1 0 0 0 0 0 0 0
0 1 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0`)
		fmt.Printf("%q\n", board)
		if valid, err := board.Valid(); valid && err != nil {
			t.Errorf("expected duplicate number error")
		}
	})
	t.Run("duplicate in region 3", func(t *testing.T) {
		board, _ := NewBoard(`
0 0 0 0 0 0 1 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 1
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0`)
		fmt.Printf("%q\n", board)
		if valid, err := board.Valid(); valid && err != nil {
			t.Errorf("expected duplicate number error")
		}
	})
}

func Test_Hardest(t *testing.T) {
	const (
		hardest = `
+-------+-------+-------+
| 8 0 0 | 0 0 0 | 0 0 0 |
| 0 0 3 | 6 0 0 | 0 0 0 |
| 0 7 0 | 0 9 0 | 2 0 0 |
+-------+-------+-------+
| 0 5 0 | 0 0 7 | 0 0 0 |
| 0 0 0 | 0 4 5 | 7 0 0 |
| 0 0 0 | 1 0 0 | 0 3 0 |
+-------+-------+-------+
| 0 0 1 | 0 0 0 | 0 6 8 |
| 0 0 8 | 5 0 0 | 0 1 0 |
| 0 9 0 | 0 0 0 | 4 0 0 |
+-------+-------+-------+`
		hardestSolution = `
+-------+-------+-------+
| 8 1 2 | 7 5 3 | 6 4 9 |
| 9 4 3 | 6 8 2 | 1 7 5 |
| 6 7 5 | 4 9 1 | 2 8 3 |
+-------+-------+-------+
| 1 5 4 | 2 3 7 | 8 9 6 |
| 3 6 9 | 8 4 5 | 7 2 1 |
| 2 8 7 | 1 6 9 | 5 3 4 |
+-------+-------+-------+
| 5 2 1 | 9 7 4 | 3 6 8 |
| 4 3 8 | 5 2 6 | 9 1 7 |
| 7 9 6 | 3 1 8 | 4 5 2 |
+-------+-------+-------+`
	)
	t.Run("solve", func(t *testing.T) {
		board, _ := NewBoard(hardest)
		solution, _ := NewBoard(hardestSolution)
		got, _ := board.Solve()
		if !reflect.DeepEqual(got, solution) {
			t.Errorf("solve() got %v, want %v", got, solution)
		}
	})
}

func Benchmark_Solve(b *testing.B) {
	// Benchmark_Solve-8   	   57832	     17580 ns/op
	for n := 0; n < b.N; n++ {
		board, _ := NewBoard(Problem1)
		board.Solve()
	}
}
