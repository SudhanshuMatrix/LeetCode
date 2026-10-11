import "container/heap"

type Project struct {
	capital int
	profit  int
}

type MinHeap []Project

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].capital < h[j].capital }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	heap.Init(minHeap)
	heap.Init(maxHeap)

	for i := range profits {
		heap.Push(minHeap, Project{capital[i], profits[i]})
	}

	for i := 0; i < k; i++ {
		for minHeap.Len() > 0 && (*minHeap)[0].capital <= w {
			p := heap.Pop(minHeap).(Project)
			heap.Push(maxHeap, p.profit)
		}

		if maxHeap.Len() == 0 {
			break
		}

		w += heap.Pop(maxHeap).(int)
	}

	return w
}