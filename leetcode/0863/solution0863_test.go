package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question struct {
	input
	answer
}

type input struct {
	root   []int
	target []int
	K      int
}

type answer struct {
	one []int
}

func Test_Problem0863(t *testing.T) {

	qs := []question{
		{
			input{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, 2},
			answer{[]int{7, 4, 1}},
		},
	}

	for _, q := range qs {
		tree := structures.Ints2TreeNode(q.input.root)
		target := structures.Ints2TreeNode(q.input.target)

		result := distanceK(tree, target, q.input.K)
		pass := reflect.DeepEqual(q.answer.one, result)

		fmt.Printf("【pass】:%v\t", pass)
		fmt.Printf("【input】:%v\t", q.input)
		fmt.Printf("【expect】:%v\t", q.answer.one)
		fmt.Printf("【output】:%v\n\n", result)

		if !pass {
			t.Errorf("\n input: %v \n output: %v \n expect: %v", q.input, result, q.answer.one)
		}
	}
}
