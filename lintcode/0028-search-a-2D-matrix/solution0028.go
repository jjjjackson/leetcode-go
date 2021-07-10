package lintcode

// https://www.lintcode.com/problem/28/?_from=collection&fromId=195
// 經典的 2D binary search

/**
 * @param matrix: matrix, a list of lists of integers
 * @param target: An integer
 * @return: a boolean, indicate whether matrix contains target
 */
func searchMatrix(matrix [][]int, target int) bool {

	var up, down int = 0, len(matrix) - 1
	var locatedRow int

	for up+1 < down {
		var mid int = up + (down-up)/2

		if matrix[mid][0] == target {
			return true
		}

		if matrix[mid][0] < target {
			up = mid
		} else {
			down = mid
		}
	}

	if matrix[down][0] <= target {
		locatedRow = down
	} else if matrix[up][0] <= target {
		locatedRow = up
	} else {
		return false
	}

	var left, right int = 0, len(matrix[0]) - 1

	for left+1 < right {
		var mid int = left + (right-left)/2

		if matrix[locatedRow][mid] == target {
			return true
		} else if matrix[locatedRow][mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if matrix[locatedRow][left] == target || matrix[locatedRow][right] == target {
		return true
	}

	return false
}
