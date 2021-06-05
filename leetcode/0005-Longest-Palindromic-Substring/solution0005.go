package leetcode

// 找擁有最長迴文的子字串
// Time O(n^2)
// Space O(1)
// 從中間擴散
// 找長度偶數和奇數的最長值

// pseudo code
// for point in string.length {
// 	odd = maxPalindrome(string, point, point)
// 	even = maxPalindrome(string, point, point+1)
// 	result = max( odd, even ,result)
// }
// return result

// 可以用 Manacher's algorithm 解 Time O(n) Space(n) 的解法，但就要直接背

func longestPalindrome(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		oddStr := maxPalindrome(s, i, i)    // 奇數
		evenStr := maxPalindrome(s, i, i+1) // 偶數
		largeStr := longerString(oddStr, evenStr)
		result = longerString(largeStr, result)
	}
	return result
}

func maxPalindrome(s string, left, right int) string {
	subStr := ""

	for left >= 0 && right < len(s) && s[left] == s[right] {
		subStr = s[left : right+1]
		left--
		right++
	}

	return subStr
}

func longerString(a, b string) string {
	if len(a) >= len(b) {
		return a
	}
	return b
}
