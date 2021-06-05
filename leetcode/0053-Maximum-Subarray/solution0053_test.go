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
}

type answer struct {
	one int
}

const MAX_INT = int(1<<32 - 1)
const MIN_INT = int(-1 << 31)
const MAX_LENGTH = 30000

func Test_Problem0053(t *testing.T) {

	maxLengthArray := make([]int, MAX_LENGTH)
	for i := 0; i < MAX_LENGTH; i++ {
		maxLengthArray[i] = MAX_INT
	}

	qs := []question{
		{
			input{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			answer{6},
		},
		{
			input{[]int{1}},
			answer{1},
		},
		{
			input{[]int{5, 4, -1, 7, 8}},
			answer{23},
		},
		{
			input{[]int{1, 2}},
			answer{3},
		},
		{
			input{[]int{1, -2}},
			answer{1},
		},
		{
			input{[]int{-2, 1}},
			answer{1},
		},
		{
			input{[]int{-2, -1}},
			answer{-1},
		},
		{
			input{[]int{-1, -2}},
			answer{-1},
		},
		{
			input{[]int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}},
			answer{6},
		},
		{
			input{[]int{MAX_INT, MAX_INT}},
			answer{MAX_INT * 2},
		},
		{
			input{maxLengthArray},
			answer{MAX_INT * MAX_LENGTH},
		},
	}

	for _, q := range qs {

		result := maxSubArray(q.input.nums)
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
