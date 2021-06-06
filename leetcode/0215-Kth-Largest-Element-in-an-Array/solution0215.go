package leetcode

// https://leetcode.com/problems/kth-largest-element-in-an-array/
// 思路 => 最大值的 priority queue
// Time O(n * log(n)) ->  從 Heap Pop 值
// Space O(n) -> Heap 的大小

import "container/heap"

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a] > pq[b]
}

func (pq PriorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	r := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return r
}

func findKthLargest(nums []int, k int) int {
	pq := make(PriorityQueue, len(nums))

	for i, v := range nums {
		pq[i] = v
	}

	heap.Init(&pq)

	result := 0
	for i := 0; i < k; i++ {
		result = heap.Pop(&pq).(int)
	}

	return result
}
