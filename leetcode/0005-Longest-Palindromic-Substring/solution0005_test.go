package leetcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	s string
}

type Answer []string

const MAX_INT = int(1<<32 - 1)
const MIN_INT = int(-1 << 31)
const MAX_LENGTH = 30000

func Test_Solution0005(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		t.Helper()

		result := longestPalindrome(q.Input.s)
		pass := contains(q.Answer, result)

		if !pass {
			t.Errorf("\n Input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		q := &Question{Input{"babad"}, []string{"aba", "bab"}}
		checker(t, q)
	})

	t.Run("Boundary", func(t *testing.T) {

		maxLengthArray := make([]int, MAX_LENGTH)
		for i := 0; i < MAX_LENGTH; i++ {
			maxLengthArray[i] = MAX_INT
		}

		qs := []*Question{
			{
				Input{"a"},
				[]string{"a"},
			},
			{
				Input{"aa"},
				[]string{"aa"},
			},
			{
				Input{"aaa"},
				[]string{"aaa"},
			},
			{
				Input{"ac"},
				[]string{"a", "c"},
			},
			{
				Input{"cbbd"},
				[]string{"bb"},
			},
			{
				Input{"ccdd"},
				[]string{"cc", "dd"},
			},
			{
				Input{""},
				[]string{""},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
