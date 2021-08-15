package leetcode

//https://leetcode.com/problems/binary-tree-postorder-traversal/

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, result)
	traversal(root.Right, result)
	*result = append(*result, root.Val)

	return
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}

	traversal(root, &result)
	return result
}
