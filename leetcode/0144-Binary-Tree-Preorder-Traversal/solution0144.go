package leetcode

// https://leetcode.com/problems/binary-tree-preorder-traversal/

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

// For loo 的寫法
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		result = append(result, node.Val)
		stack = append(stack, node.Right) // 注意這個地方的順序
		stack = append(stack, node.Left)
	}

	return result
}

// Recursive 的寫法
// func traversal(root *TreeNode, result *[]int) {

// 	if root == nil {
// 		return
// 	}

// 	*result = append(*result, root.Val)
// 	traversal(root.Left, result)
// 	traversal(root.Right, result)
// }

// func preorderTraversal(root *TreeNode) []int {
// 	result := []int{}

// 	if root == nil {
// 		return []int{}
// 	}
// 	traversal(root, &result)
// 	return result
// }
