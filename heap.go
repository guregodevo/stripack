package stripack


func Array(dx, dy int) [][]int {
    pic := make([][]int, dx) /* type declaration */
    for i := range pic {
        pic[i] = make([]int, dy) /* again the type? */
        for j := range pic[i] {
            pic[i][j] = -1
        }
    }
    return pic
}

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
    //item.index = -1 // for safety
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
    //item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

