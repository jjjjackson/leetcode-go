package leetcode

import (
	"testing"
)

type question struct {
	input
	answer
}

type input struct {
	benginWord string
	endWord    string
	wordList   []string
}

type answer struct {
	one int
}

func Test_Problem(t *testing.T) {

	qs := []question{
		{
			input{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			answer{5},
		},
		{
			input{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			answer{0},
		},
	}

	for _, q := range qs {

		result := ladderLength(q.input.benginWord, q.input.endWord, q.wordList)
		pass := q.answer.one == result

		if !pass {
			t.Errorf("\n input: %v \n output: %v \n expect: %v", q.input, result, q.answer.one)
		}
	}
}
