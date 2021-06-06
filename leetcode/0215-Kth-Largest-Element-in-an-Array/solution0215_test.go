package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	input
	answer
}

type input struct {
	nums []int
	k    int
}

type answer struct {
	one int
}

// 1 <= k <= nums.length <= 10^4
// -104 <= nums[i] <= 10^4

const MAX_INT = 10000
const MIN_INT = -10000
const MAX_LENGTH = 10000

func Test_Problem0215(t *testing.T) {

	maxLengthArray := make([]int, MAX_LENGTH)
	for i := 0; i < MAX_LENGTH; i++ {
		maxLengthArray[i] = MAX_INT
	}

	qs := []question{
		{
			input{[]int{3, 2, 1, 5, 6, 4}, 2},
			answer{5},
		},
		{
			input{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			answer{4},
		},
		{
			input{[]int{1}, 1},
			answer{1},
		},
		{
			input{[]int{1, 2}, 1},
			answer{2},
		},
		{
			input{[]int{MIN_INT, 0}, 1},
			answer{0},
		},
		{
			input{[]int{MAX_INT, MIN_INT}, 1},
			answer{MAX_INT},
		},
	}

	for _, q := range qs {

		result := findKthLargest(q.input.nums, q.input.k)
		pass := q.answer.one == result

		fmt.Printf("【pass】:%v\t", pass)
		fmt.Printf("【input】:%v\t", q.input)
		fmt.Printf("【expect】:%v\t", q.answer.one)
		fmt.Printf("【output】:%v\n\n", result)

		if !pass {
			t.Errorf("\n input: %v \n output: %v \n expect: %v", q.input, result, q.answer.one)
		}
	}
}
