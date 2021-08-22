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

type Answer = [][]int

func Test_Solution0094(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		root := structures.Ints2TreeNode(q.Input.Root)

		result := levelOrder(root)
		pass := reflect.DeepEqual(result, q.Answer)

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
				[][]int{{3}, {9, 20}, {15, 7}},
			},
			{
				Input{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}},
				[][]int{{1}, {2, 3}, {4, 5}},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

}
