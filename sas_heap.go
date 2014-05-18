package stripack

type NarrowRectHeap []*Rect

func (h NarrowRectHeap) Len() int           { return len(h) }
func (h NarrowRectHeap) Less(i, j int) bool { return h[i].H > h[j].H }
func (h NarrowRectHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (pq *NarrowRectHeap) Push(x interface{}) {
    *pq = append(*pq, x.(*Rect))
}

func (pq *NarrowRectHeap) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

type WideRectHeap []*Rect

func (h WideRectHeap) Len() int           { return len(h) }
func (h WideRectHeap) Less(i, j int) bool { return h[i].W > h[j].W }
func (h WideRectHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (pq *WideRectHeap) Push(x interface{}) {
    *pq = append(*pq, x.(*Rect))
}

func (pq *WideRectHeap) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

