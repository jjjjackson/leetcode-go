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

func Test_Solution0094(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		tree := structures.Ints2TreeNode(q.Input.Root)

		result := inorderTraversal(tree)
		pass := reflect.DeepEqual(q.Answer, result)

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{1, structures.NULL, 2, structures.NULL, 3}},
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
				Input{[]int{1, structures.NULL, 2}},
				[]int{1, 2},
			},
			{
				Input{[]int{1, 2, 3, structures.NULL, structures.NULL, 4}},
				[]int{2, 1, 4, 3},
			},
			{
				Input{[]int{1, 2, 3, structures.NULL, structures.NULL, 4}},
				[]int{2, 1, 4, 3},
			},
			{
				Input{[]int{1, 2, 3, 4, 5, 6, 7, structures.NULL, structures.NULL, 8, 9, structures.NULL, structures.NULL, structures.NULL, structures.NULL}},
				[]int{4, 2, 8, 5, 9, 1, 6, 3, 7},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

}
