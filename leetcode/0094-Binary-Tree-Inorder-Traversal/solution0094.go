package leetcode

// https://leetcode.com/problems/binary-tree-inorder-traversal/

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

func traversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	result := []int{}
	if root.Left != nil {
		result = traversal(root.Left)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		rightResult := traversal(root.Right)
		result = append(result, rightResult...)
	}

	return result
}

func inorderTraversal(root *TreeNode) []int {
	result := traversal(root)

	return result
}
