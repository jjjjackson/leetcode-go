package leetcode

import (
	"reflect"
	"testing"
)

type question struct {
	input
	expect
}

type input struct {
	grid [][]byte
}

type expect = [][]byte

func Test_Problem0200(t *testing.T) {

	checker := func(t testing.TB, q *question) {
		t.Helper()

		solve(q.input.grid)
		pass := reflect.DeepEqual(q.expect, q.input.grid)
		if !pass {
			t.Errorf("\n output: %v \n expect: %v \n", q.input.grid, q.expect)
		}
	}

	t.Run("Given by question 1", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}},
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		}

		checker(t, q)
	})

	t.Run("Given by question 2", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}},
			[][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}},
		}

		checker(t, q)
	})

	t.Run("Given by question 3", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}}},
			[][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}},
		}

		checker(t, q)
	})

	t.Run("Given by question 4", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}},
			[][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}},
		}

		checker(t, q)
	})

	t.Run("TestBoundary", func(t *testing.T) {
		qs := []*question{
			{
				input{[][]byte{{'X'}}},
				[][]byte{{'X'}},
			},
			{
				input{[][]byte{{'O'}}},
				[][]byte{{'O'}},
			},
			{
				input{[][]byte{{'X', 'X'}}},
				[][]byte{{'X', 'X'}},
			},
			{
				input{[][]byte{{'O', 'O'}}},
				[][]byte{{'O', 'O'}},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

	t.Run("X Quantain By O", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'O', 'O', 'O'}, {'O', 'X', 'O'}, {'O', 'O', 'O'}}},
			[][]byte{{'O', 'O', 'O'}, {'O', 'X', 'O'}, {'O', 'O', 'O'}},
		}

		checker(t, q)
	})

	t.Run("O Quantain By X in 4*4", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'O', 'O', 'O', 'O'}, {'O', 'X', 'X', 'O'}, {'O', 'X', 'X', 'O'}, {'O', 'O', 'O', 'O'}}},
			[][]byte{{'O', 'O', 'O', 'O'}, {'O', 'X', 'X', 'O'}, {'O', 'X', 'X', 'O'}, {'O', 'O', 'O', 'O'}},
		}

		checker(t, q)
	})

	t.Run("X Quantain By O in 4*4", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'X', 'X'}}},
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}},
		}

		checker(t, q)
	})

	t.Run("X Quantain By O But have break point in 4*4", func(t *testing.T) {
		q := &question{
			input{[][]byte{{'X', 'O', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'X', 'X'}}},
			[][]byte{{'X', 'O', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'X', 'X'}},
		}

		checker(t, q)
	})
}
