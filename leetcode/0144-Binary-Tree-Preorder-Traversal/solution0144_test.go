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
}

type Answer = []int

func Test_Solution0144(t *testing.T) {
	checker := func(t testing.TB, q *Question) {
		tree := structures.Ints2TreeNode(q.Input.Root)

		result := preorderTraversal(tree)
		pass := reflect.DeepEqual(q.Answer, result)

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{1, structures.NULL, 2, 3}},
				[]int{1, 2, 3},
			},
			{
				Input{[]int{structures.NULL}},
				[]int{structures.NULL},
			},
			{
				Input{[]int{1}},
				[]int{1},
			},
			{
				Input{[]int{1, 2}},
				[]int{1, 2},
			},
			{
				Input{[]int{1, structures.NULL, 2}},
				[]int{1, 2},
			},
			{
				Input{[]int{1, 2, 4, 3, structures.NULL, 5, structures.NULL}},
				[]int{1, 2, 3, 4, 5},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
