package leetcode

import (
	"testing"
)

type question struct {
	input
	answer
}

type input struct {
	s string
}

type answer struct {
	one []string
}

const MAX_INT = int(1<<32 - 1)
const MIN_INT = int(-1 << 31)
const MAX_LENGTH = 30000

func Test_Problem0005(t *testing.T) {

	maxLengthArray := make([]int, MAX_LENGTH)
	for i := 0; i < MAX_LENGTH; i++ {
		maxLengthArray[i] = MAX_INT
	}

	qs := []question{
		{
			input{"babad"},
			answer{[]string{"aba", "bab"}},
		},
		{
			input{"a"},
			answer{[]string{"a"}},
		},
		{
			input{"aa"},
			answer{[]string{"aa"}},
		},
		{
			input{"aaa"},
			answer{[]string{"aaa"}},
		},
		{
			input{"ac"},
			answer{[]string{"a", "c"}},
		},
		{
			input{"cbbd"},
			answer{[]string{"bb"}},
		},
		{
			input{"ccdd"},
			answer{[]string{"cc", "dd"}},
		},
		{
			input{""},
			answer{[]string{""}},
		},
	}

	for _, q := range qs {

		result := longestPalindrome(q.input.s)
		pass := contains(q.answer.one, result)

		if !pass {
			t.Errorf("\n input: %v \n output: %v \n expect: %v", q.input, result, q.answer.one)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
