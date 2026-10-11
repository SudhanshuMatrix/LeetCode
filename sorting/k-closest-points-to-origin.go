package main

import "container/heap"

type Pair struct {
	point []int
	dist  int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, point := range points {
		dist := point[0]*point[0] + point[1]*point[1]

		heap.Push(h, Pair{point, dist})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([][]int, k)

	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(h).(Pair).point
	}

	return ans
}