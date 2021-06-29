package leetcode

import (
	"testing"

	"github.com/jjjjackson/leetcode-go/helper"
)

type Question struct {
	input
	expect
}

type input struct {
	digits string
}

type expect = []string

func Test_Solution0017(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		t.Helper()

		got := letterCombinations(q.input.digits)

		pass, err := helper.SliceEqual(q.expect, got)
		if err != nil {
			t.Fatalf("Error: incorrect slice comparsion")
		}

		if !pass {
			t.Errorf("\ninputs: %v \n output: %v \n expect: %v \n", q.input.digits, got, q.expect)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		q := &Question{
			input{"23"},
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		}

		checker(t, q)
	})

	t.Run("Boundary", func(t *testing.T) {
		qs := []*Question{
			{
				input{""},
				[]string{},
			},
			{
				input{"2"},
				[]string{"a", "b", "c"},
			},
			{
				input{"7"},
				[]string{"p", "q", "r", "s"},
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
