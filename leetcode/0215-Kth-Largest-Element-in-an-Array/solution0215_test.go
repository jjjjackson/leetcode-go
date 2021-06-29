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
	K    int
}

type Answer = int

// 1 <= k <= nums.length <= 10^4
// -104 <= nums[i] <= 10^4

const MAX_INT = 10000
const MIN_INT = -10000
const MAX_LENGTH = 10000

func Test_Solution0215(t *testing.T) {
	checker := func(t testing.TB, q *Question) {

		result := findKthLargest(q.Input.Nums, q.Input.K)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[]int{3, 2, 1, 5, 6, 4}, 2},
				5,
			},
			{
				Input{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
				4,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

	t.Run("Boundary", func(t *testing.T) {
		maxLengthArray := make([]int, MAX_LENGTH)
		for i := 0; i < MAX_LENGTH; i++ {
			maxLengthArray[i] = MAX_INT
		}

		qs := []*Question{
			{
				Input{[]int{1}, 1},
				1,
			},
			{
				Input{[]int{1, 2}, 1},
				2,
			},
			{
				Input{[]int{MIN_INT, 0}, 1},
				0,
			},
			{
				Input{[]int{MAX_INT, MIN_INT}, 1},
				MAX_INT,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})

}
