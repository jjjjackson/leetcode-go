package leetcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	board [][]byte
	word  string
}

type Answer = bool

func Test_Solution0079(t *testing.T) {
	checker := func(t testing.TB, q *Question) {
		result := exist(q.Input.board, q.word)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n input: %q \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"},
				true,
			},
			{
				Input{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"},
				true,
			},
			{
				Input{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"},
				false,
			},
			{
				Input{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCFB"},
				false,
			},
			{
				Input{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS"},
				true,
			},
			{
				Input{[][]byte{{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}}, "aaaaaaaaaaaaa"},
				false,
			},
			{
				Input{
					[][]byte{{'a', 'a', 'b', 'a', 'a', 'b'}, {'a', 'a', 'b', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a', 'b', 'a'}, {'b', 'a', 'b', 'b', 'a', 'b'}, {'a', 'b', 'b', 'a', 'b', 'a'}, {'b', 'a', 'a', 'a', 'a', 'b'}},
					"bbbaabbbbbab",
				},
				false,
			},
			{
				Input{
					[][]byte{{'a'}},
					"a",
				},
				true,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
