package leetcode

import (
	"testing"
)

type Question struct {
	Input
	Answer
}

type Input struct {
	benginWord string
	endWord    string
	wordList   []string
}

type Answer = int

func Test_Solution0127(t *testing.T) {

	checker := func(t testing.TB, q *Question) {
		result := ladderLength(q.Input.benginWord, q.Input.endWord, q.wordList)
		pass := q.Answer == result

		if !pass {
			t.Errorf("\n input: %v \n output: %v \n expect: %v", q.Input, result, q.Answer)
		}
	}

	t.Run("Given by question", func(t *testing.T) {
		qs := []*Question{
			{
				Input{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
				5,
			},
			{
				Input{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
				0,
			},
		}

		for _, q := range qs {
			checker(t, q)
		}
	})
}
