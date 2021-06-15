package leetcode

// https://leetcode.com/problems/word-ladder/description/
// BFS
// 先找出只相差一個字的字群，
// 用 Map 等級已經走過的字串
// 然後用 BFS 找最近距離
// 優化：可以用 DP 把搜尋過的鄰近字存起來（用List之類的，組成 Graph）

func ladderLength(beginWord string, endWord string, wordList []string) int {

	visited := make(map[string]bool)
	queue := []string{beginWord}
	layer := 0

	for len(queue) > 0 {
		layer++
		size := len(queue)

		for size > 0 {
			size--
			w := queue[0]
			queue = queue[1:]

			if w == endWord {
				return layer
			}

			if _, ok := visited[w]; ok {
				continue
			}

			visited[w] = true
			neighbor := getRelativeWord(wordList, w)

			for _, n := range neighbor {
				if _, ok := visited[n]; ok {
					continue
				}
				queue = append(queue, n)
			}
		}
	}

	return 0
}

func getRelativeWord(wordList []string, word string) []string {
	result := []string{}

	for _, w := range wordList {
		diffCount := 0
		if w == word {
			continue
		}
		for i := range w {

			if w[i] != word[i] {
				diffCount++
			}
		}

		if diffCount == 1 {
			result = append(result, w)
		}
	}

	return result
}
