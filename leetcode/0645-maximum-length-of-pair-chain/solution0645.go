package leetcode

import "sort"

// https://leetcode.com/problems/maximum-length-of-pair-chain
// DP => 但我不懂是怎麼做的

const MIN_INT = int(-1 << 31)

func findLongestChain(pairs [][]int) int {

	if len(pairs) <= 1 {
		return len(pairs)
	}

	sort.Slice(pairs[:], func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	dp := make([]int, len(pairs))
	for i := 0; i < len(pairs); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(pairs); i++ {
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	maxV := 0
	for _, v := range dp {
		maxV = max(maxV, v)
	}

	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
