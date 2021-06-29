package leetcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	grid [][]byte
}

type Answer = int

func Test_Solution0200(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		result := numIslands(q.Input.grid)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				}},
				1,
			},
			{
				Input{[][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				}},
				3,
			},
			{
				Input{[][]byte{
					{'0', '0', '0', '0', '0'},
				}},
				0,
			},
		}
		for _, q := range qs {
			checker(t, q)
		}
	})
}
