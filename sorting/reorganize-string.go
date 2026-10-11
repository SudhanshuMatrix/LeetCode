type Pair struct {
	ch    byte
	count int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func reorganizeString(s string) string {
	cnt := make([]int, 26)

	for _, c := range s {
		cnt[c-'a']++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			heap.Push(h, Pair{byte('a' + i), cnt[i]})
		}
	}

	ans := make([]byte, 0, len(s))

	var prev Pair

	for h.Len() > 0 {
		cur := heap.Pop(h).(Pair)

		if prev.count > 0 && cur.ch == prev.ch {
			if h.Len() == 0 {
				return ""
			}

			next := heap.Pop(h).(Pair)

			ans = append(ans, next.ch)
			next.count--

			if next.count > 0 {
				heap.Push(h, next)
			}

			heap.Push(h, cur)
			prev = next
		} else {
			ans = append(ans, cur.ch)
			cur.count--

			if cur.count > 0 {
				heap.Push(h, cur)
			}

			prev = cur
		}
	}

	return string(ans)
}