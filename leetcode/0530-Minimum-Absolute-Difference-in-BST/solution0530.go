package leetcode

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const MAX_INT = int(1<<32 - 1)

func inorder(root *TreeNode, prevNum, minDiff *int) {
	if root == nil {
		return
	}

	inorder(root.Left, prevNum, minDiff)

	if *prevNum != -1 {
		*minDiff = min(*minDiff, abs(root.Val-*prevNum))
	}

	*prevNum = root.Val
	inorder(root.Right, prevNum, minDiff)
}

func getMinimumDifference(root *TreeNode) int {
	prevNum, minDiff := -1, MAX_INT
	inorder(root, &prevNum, &minDiff)
	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
