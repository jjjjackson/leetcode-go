package leetcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	IsConnected [][]int
}

type Answer = int

func Test_Solution0547(t *testing.T) {

	checker := func(t testing.TB, q *Question) {

		result := findCircleNum(q.Input.IsConnected)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
				2,
			},
			{
				Input{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
				3,
			},
			{
				Input{[][]int{{1, 0}, {0, 1}}},
				2,
			},
			{
				Input{[][]int{{1, 1}, {1, 1}}},
				1,
			},
		}
		for _, q := range qs {
			checker(t, q)
		}
	})
}
