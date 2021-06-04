package leetcode

// 思路 =>  貪心算法
// 只找最大區間
// 如果 sum < 0 就放棄那個 Sum 重新尋找新區間
// pseudo
// max = min_int
// sum = 0
// loop n in nums {
// 	sum += n
// 	max = Math.max(max, sum)
// 	sum = Math.max(sum,0)
// }
// return max

func maxSubArray(nums []int) int {
	const MIN_INT = int(-1 << 31)

	maxSum, sum := MIN_INT, 0

	for _, n := range nums {
		sum += n
		maxSum = max(maxSum, sum)
		sum = max(sum, 0)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
