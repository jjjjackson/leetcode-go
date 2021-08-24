package leetcode

// https://leetcode.com/problems/validate-binary-search-tree/submissions/
// 題目的邊界很 Tricky

import (
	"math"

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

func isValidBST(root *TreeNode) bool {
	return isValidNodeInBST(root, math.Inf(-1), math.Inf(1))
}

func isValidNodeInBST(root *TreeNode, minVal, maxVal float64) bool {
	if root == nil {
		return true
	}

	value := float64(root.Val)
	if value <= minVal || value >= maxVal {
		return false
	}

	return isValidNodeInBST(root.Left, minVal, value) && isValidNodeInBST(root.Right, value, maxVal)
}
