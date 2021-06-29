package leetcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	Nums []int
}

type Answer = int

const MAX_INT = int(1<<32 - 1)
const MIN_INT = int(-1 << 31)
const MAX_LENGTH = 30000

func Test_Solution0053(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		result := maxSubArray(q.Input.Nums)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		q := &Question{
			Input{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			6,
		}

		checker(t, q)
	})

	t.Run("Boundary", func(t *testing.T) {
		maxLengthArray := make([]int, MAX_LENGTH)
		for i := 0; i < MAX_LENGTH; i++ {
			maxLengthArray[i] = MAX_INT
		}

		qs := []*Question{
			{
				Input{[]int{1}},
				1,
			},
			{
				Input{[]int{5, 4, -1, 7, 8}},
				23,
			},
			{
				Input{[]int{1, 2}},
				3,
			},
			{
				Input{[]int{1, -2}},
				1,
			},
			{
				Input{[]int{-2, 1}},
				1,
			},
			{
				Input{[]int{-2, -1}},
				-1,
			},
			{
				Input{[]int{-1, -2}},
				-1,
			},
			{
				Input{[]int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}},
				6,
			},
			{
				Input{[]int{MAX_INT, MAX_INT}},
				MAX_INT * 2,
			},
			{
				Input{maxLengthArray},
				MAX_INT * MAX_LENGTH,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
