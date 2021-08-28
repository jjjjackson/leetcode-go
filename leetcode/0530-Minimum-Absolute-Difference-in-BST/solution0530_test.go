package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	Root []int
}

type Answer = int

func Test_Solution0530(t *testing.T) {
	checker := func(t testing.TB, q *Question) {
		tree := structures.Ints2TreeNode(q.Input.Root)

		result := getMinimumDifference(tree)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{4, 2, 6, 1, 3}},
				1,
			},
			{
				Input{[]int{236, 104, 701, structures.NULL, 227, structures.NULL, 911}},
				9,
			},
			{
				Input{[]int{1, 0, 48, structures.NULL, structures.NULL, 12, 49}},
				1,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
