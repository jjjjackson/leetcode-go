package leetcode

import (
	"reflect"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	Root []int
	Val  int
}

type Answer = []int

func Test_Solution0700(t *testing.T) {
	checker := func(t testing.TB, q *Question) {
		tree := structures.Ints2TreeNode(q.Input.Root)

		result := searchBST(tree, q.Input.Val)
		resultInts := structures.Tree2ints(result)
		pass := reflect.DeepEqual(q.Answer, resultInts)

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{4, 2, 7, 1, 3}, 2},
				[]int{2, 1, 3},
			},
			{
				Input{[]int{4, 2, 7, 1, 3}, 5},
				[]int{},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
