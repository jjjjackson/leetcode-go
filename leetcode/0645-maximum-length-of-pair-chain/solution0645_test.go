package leetcode

import "testing"

type Question struct {
	Input
	Expect
}

type Input struct {
	Pairs [][]int
}

type Expect = int

func Test_Solution0645(t *testing.T) {
	checker := func(t testing.TB, q *Question) {
		t.Helper()

		got := findLongestChain(q.Input.Pairs)
		pass := got == q.Expect

		if !pass {
			t.Errorf("\ninputs: %v \n output: %v \n expect: %v \n", q.Input.Pairs, got, q.Expect)
		}
	}
	t.Run("Given by question 1", func(t *testing.T) {
		q := Question{Input{[][]int{{1, 2}, {2, 3}, {3, 4}}}, 2}
		checker(t, &q)
	})
	t.Run("Given by question 2", func(t *testing.T) {
		q := Question{Input{[][]int{{1, 2}, {7, 8}, {4, 5}}}, 3}
		checker(t, &q)
	})
	t.Run("Given by question 3", func(t *testing.T) {
		q := Question{Input{[][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}}}, 4}
		checker(t, &q)
	})

}
