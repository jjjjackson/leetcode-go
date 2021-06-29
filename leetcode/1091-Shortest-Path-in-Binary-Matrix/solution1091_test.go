package leetcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	grid [][]int
}

type Answer = int

func Test_Solution1091(t *testing.T) {
	checker := func(t testing.TB, q *Question) {
		result := shortestPathBinaryMatrix(q.Input.grid)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {

		qs := []*Question{
			{
				Input{[][]int{{0, 1}, {1, 0}}},
				2,
			},
			{
				Input{[][]int{{0, 0}, {0, 1}}},
				-1,
			},
			{
				Input{[][]int{{1, 0}, {0, 0}}},
				-1,
			},
			{
				Input{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}},
				4,
			},
			{
				Input{[][]int{{0, 0, 0, 1}, {0, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}}},
				4,
			},
			{
				Input{[][]int{{0, 0, 0}, {1, 1, 1}, {1, 1, 0}}},
				-1,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
