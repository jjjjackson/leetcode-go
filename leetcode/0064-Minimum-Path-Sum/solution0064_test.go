package leetcode

import (
	"testing"
)

type question struct {
	input
	expect
}

type input struct {
	grid [][]int
}

type expect = int

func Test_Solution0064(t *testing.T) {

	checker := func(t testing.TB, q *question) {
		t.Helper()

		got := minPathSum(q.input.grid)
		pass := got == q.expect

		if !pass {
			t.Errorf("\ninputs: %v \n output: %v \n expect: %v \n", q.input.grid, got, q.expect)
		}
	}

	t.Run("Given by question 1", func(t *testing.T) {
		q := &question{
			input{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}},
			7,
		}

		checker(t, q)
	})

	t.Run("Given by question 2", func(t *testing.T) {
		q := &question{
			input{[][]int{{1, 2, 3}, {4, 5, 6}}},
			12,
		}

		checker(t, q)
	})

	t.Run("Given by question 3", func(t *testing.T) {
		q := &question{
			input{[][]int{{1, 2}, {1, 1}}},
			3,
		}

		checker(t, q)
	})

	t.Run("Boundary", func(t *testing.T) {
		qs := []*question{
			{
				input{[][]int{{}}},
				0,
			},
			{
				input{[][]int{{1}}},
				1,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
