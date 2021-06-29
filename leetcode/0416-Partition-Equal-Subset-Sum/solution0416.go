package leetcode

// https://leetcode.com/problems/partition-equal-subset-sum/description/
// DP 背包問題
// 看能不能放滿 sum / 2

func canPartition(nums []int) bool {

	sumOfArr := getSumOfArray(nums)
	if sumOfArr%2 != 0 {
		return false
	}

	var availability int = sumOfArr / 2
	dp := make([]bool, availability+1)
	dp[0] = true

	for _, num := range nums {
		for i := availability; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[availability]
}

func getSumOfArray(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}
