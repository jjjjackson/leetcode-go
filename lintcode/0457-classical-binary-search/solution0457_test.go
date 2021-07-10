package lintcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	nums   []int
	target int
}

type Answer []int

func Test_Solution0457(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		t.Helper()

		result := findPosition(q.Input.nums, q.Input.target)
		pass := contains(q.Answer, result)

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{Input{[]int{1, 2, 2, 4, 5, 5}, 2}, []int{1, 2}},
			{Input{[]int{1, 2, 2, 4, 5, 5}, 6}, []int{-1}},
			{Input{[]int{1, 2, 2, 4, 6, 6}, 5}, []int{-1}},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
