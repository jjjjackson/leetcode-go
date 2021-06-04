package leetcode

// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

import (
	"fmt"

	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

// 思路：
// DFS 找 Parent -> BFS 找距離

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := make(map[*TreeNode]*TreeNode)

	target = findTarget(root, target.Val)
	recordParents(parents, root, nil)

	queue := []*TreeNode{nil, target}
	seen := make(map[*TreeNode]bool)
	answer := []int{}

	seen[target] = true
	seen[nil] = true

	dist := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		for _, q := range queue {
			if q != nil {
				fmt.Println(q.Left, q.Right, q.Val)
			}
		}

		fmt.Println()
		for q := range seen {
			if q != nil {
				fmt.Println(q)
			}
		}

		if node == nil {
			if dist == k {
				for _, n := range queue {
					answer = append(answer, n.Val)
				}
				return answer
			}
			queue = append(queue, nil)
			dist++
		} else {

			if _, ok := seen[node.Left]; !ok {
				seen[node.Left] = true
				queue = append(queue, node.Left)
			}

			if _, ok := seen[node.Right]; !ok {
				seen[node.Right] = true
				queue = append(queue, node.Right)
			}

			p, _ := parents[node]
			if _, ok := seen[p]; !ok {
				seen[p] = true
				queue = append(queue, p)
			}
		}

	}

	return []int{}
}

func recordParents(parents map[*TreeNode]*TreeNode, node *TreeNode, parent *TreeNode) {
	if node != nil {
		parents[node] = parent
		recordParents(parents, node.Left, node)
		recordParents(parents, node.Right, node)
	}
}

func findTarget(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}

	left := findTarget(node.Left, val)
	right := findTarget(node.Right, val)
	if left != nil {
		return left
	}
	return right
}
