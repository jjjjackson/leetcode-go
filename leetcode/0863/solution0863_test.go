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
	Root   []int
	Target []int
	K      int
}

type Answer = []int

func Test_Solution0863(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		tree := structures.Ints2TreeNode(q.Input.Root)
		target := structures.Ints2TreeNode(q.Input.Target)

		result := distanceK(tree, target, q.Input.K)
		pass := reflect.DeepEqual(q.Answer, result)

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, 2},
				[]int{7, 4, 1},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

}
