package leetcode

import "testing"

type Question struct {
	Input
	Expect
}

type Input struct {
	Nums []int
}

type Expect = bool

func Test_Solution0415(t *testing.T) {
	checker := func(t testing.TB, q *Question) {
		t.Helper()

		got := canPartition(q.Input.Nums)
		pass := got == q.Expect

		if !pass {
			t.Errorf("\ninputs: %v \n output: %v \n expect: %v \n", q.Input.Nums, got, q.Expect)
		}
	}
	t.Run("Given by question 1", func(t *testing.T) {
		q := Question{Input{[]int{1, 5, 11, 5}}, true}
		checker(t, &q)
	})
	t.Run("Given by question 2", func(t *testing.T) {
		q := Question{Input{[]int{1, 2, 3, 5}}, false}
		checker(t, &q)
	})
	t.Run("Test Boundary", func(t *testing.T) {
		qs := []*Question{
			{Input{[]int{1}}, false},
			{Input{[]int{2}}, false},
			{Input{[]int{1, 2}}, false},       // odd
			{Input{[]int{1, 3}}, false},       // even
			{Input{[]int{1, 2, 3}}, true},     // even
			{Input{[]int{1, 2, 3, 8}}, false}, // even but cant find
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
