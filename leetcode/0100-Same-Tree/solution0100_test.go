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
	Origin []int
	Target []int
}

type Answer = bool

func Test_Solution0094(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		originTree := structures.Ints2TreeNode(q.Input.Origin)
		targetTree := structures.Ints2TreeNode(q.Input.Target)

		result := isSameTree(originTree, targetTree)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{1, 2, 3}, []int{1, 2, 3}},
				true,
			},
			{
				Input{[]int{1, 2, 3}, []int{1, structures.NULL, 2}},
				false,
			},
			{
				Input{[]int{1, 2, 1}, []int{1, 1, 2}},
				false,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

}
