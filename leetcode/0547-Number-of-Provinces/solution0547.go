package leetcode

// https://leetcode.com/problems/number-of-provinces/
// 把 isConnected 看成無向圖
// DFS 找出連結起來的的人

func findCircleNum(isConnected [][]int) int {
	result := 0
	visited := make([]bool, len(isConnected))

	for p := range isConnected {
		if !visited[p] {
			findConnected(isConnected, visited, p)
			result++
		}
	}

	return result
}

func findConnected(isConnected [][]int, visited []bool, from int) {
	stack := []int{from}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[p] {
			continue
		}

		visited[p] = true

		for f := range isConnected[p] {
			if isConnected[p][f] == 1 {
				stack = append(stack, f)
			}
		}

	}

}
