package leetcode

import (
	"testing"
)

type question struct {
	input
	answer
}

type input struct {
	isConnected [][]int
}

type answer struct {
	one int
}

func Test_Problem0547(t *testing.T) {

	qs := []question{
		{
			input{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
			answer{2},
		},
		{
			input{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
			answer{3},
		},
		{
			input{[][]int{{1, 0}, {0, 1}}},
			answer{2},
		},
		{
			input{[][]int{{1, 1}, {1, 1}}},
			answer{1},
		},
	}

	for _, q := range qs {

		result := findCircleNum(q.input.isConnected)
		pass := q.answer.one == result

		if !pass {
			t.Errorf("\n input: %v \n output: %v \n expect: %v", q.input, result, q.answer.one)
		}
	}
}
