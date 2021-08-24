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

type Answer = bool

func Test_Solution0098(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		tree := structures.Ints2TreeNode(q.Input.Root)

		result := isValidBST(tree)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{2, 1, 3}},
				true,
			},
			{
				Input{[]int{2, 2, 2}},
				false,
			},
			{
				Input{[]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6}},
				false,
			},
			{
				Input{[]int{5, 1, 6, structures.NULL, structures.NULL, 3, 7}},
				false,
			},
			{
				Input{[]int{-2147483648}},
				true,
			},
			{
				Input{[]int{-2147483648, structures.NULL, -2147483648}},
				false,
			},
			{
				Input{[]int{-2147483648, structures.NULL, 2147483648}},
				true,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

}
