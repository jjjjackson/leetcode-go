package leetcode

import (
	"testing"
)

type question struct {
	input
	answer
}

type input struct {
	grid [][]byte
}

type answer struct {
	one int
}

func Test_Problem0200(t *testing.T) {

	qs := []question{
		{
			input{[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			answer{1},
		},
		{
			input{[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			answer{3},
		},
		{
			input{[][]byte{
				{'0', '0', '0', '0', '0'},
			}},
			answer{0},
		},
	}

	for _, q := range qs {

		result := numIslands(q.input.grid)
		pass := q.answer.one == result

		if !pass {
			t.Errorf("\n input: %v \n output: %v \n expect: %v", q.input, result, q.answer.one)
		}
	}
}
