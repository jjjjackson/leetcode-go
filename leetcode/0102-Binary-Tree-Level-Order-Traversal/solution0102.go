package leetcode

// https://leetcode.com/problems/binary-tree-inorder-traversal/

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

//BFS
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		splitPoint := len(queue)
		currentLayerResult := make([]int, 0, splitPoint)

		for i := 0; i < splitPoint; i++ {
			node := queue[i]

			currentLayerResult = append(currentLayerResult, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[splitPoint:]
		result = append(result, currentLayerResult)
	}

	return result
}

// 自己的解法
// 每次都把 Result, current Layer和 Node 都進去
// 如果自己是 Nil 就 return
// func traversal(root *TreeNode, currentLayer int, result *map[int][]int) {
// 	if root == nil {
// 		return
// 	}

// 	if (*result)[currentLayer] == nil {
// 		(*result)[currentLayer] = []int{}
// 	}
// 	(*result)[currentLayer] = append((*result)[currentLayer], root.Val)

// 	traversal(root.Left, currentLayer+1, result)
// 	traversal(root.Right, currentLayer+1, result)
// }

// func levelOrder(root *TreeNode) [][]int {
// 	rawResult := make(map[int][]int)
// 	traversal(root, 0, &rawResult)

// 	var result [][]int
//     for i := 0 ; i < len(rawResult); i++  { // 因為 LeetCode compiler 不會過
//         if r, ok := rawResult[i]; ok {
//             result = append(result, r)
//         }
// 	}

// 	return result
// }
