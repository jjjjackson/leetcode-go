package leetcode

import (
	"errors"
	"strconv"
)

// DFS
//

var ERROR_INCORRECT_INPUT = errors.New("Incorrect Input")

var KEYBOARD = [][]string{
	{},
	{},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	words := make(map[string]bool)

	err := dfs(digits, 0, "", words)
	if err != nil {
		return []string{}
	}

	result := make([]string, len(words))

	i := 0
	for k := range words {
		result[i] = k
		i++
	}

	return result
}

func dfs(digits string, index int, result string, words map[string]bool) error {

	if index == len(digits) {
		words[result] = true
		return nil
	}

	number, err := strconv.Atoi(string(digits[index]))
	if err != nil {
		return err
	}

	if number > 9 || number < 2 {
		return ERROR_INCORRECT_INPUT
	}

	for _, v := range KEYBOARD[number] {
		dfs(digits, index+1, result+v, words)
	}

	return nil
}
