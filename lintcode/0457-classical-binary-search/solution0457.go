package lintcode

// https://www.lintcode.com/problem/457/?_from=collection&fromId=195
// 經典的 1D Binary Search

func findPosition(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1

	for left < right {
		var mid int = left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[left] {
			left = mid + 1
		}

		if nums[mid] < nums[right] {
			right = mid + 1
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}
