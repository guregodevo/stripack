package stripack


type IntSet struct {
    set map[int]struct{}
}

func NewIntSet() *IntSet {
    s := new(IntSet)
    s.set = make(map[int]struct{})
    return s
}

func (set *IntSet) Add(i int) bool {
    _, found := set.set[i]
    set.set[i] = struct{}{}
    return !found 
}

func (set *IntSet) Diff(i int) bool {
    for value, _ := range set.set {
    	if value != i {
    		return true
		}
	}
    return true 
}

func (set *IntSet) Exists(i int) bool {
    _, found := set.set[i]
    return found 
}


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

type DescHRect []*Rect

func (h DescHRect) Len() int           { return len(h) }
func (h DescHRect) Less(i, j int) bool { return h[i].H > h[j].H }
func (h DescHRect) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type DescWRect []*Rect

func (h DescWRect) Len() int           { return len(h) }
func (h DescWRect) Less(i, j int) bool { return h[i].W > h[j].W }
func (h DescWRect) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (pq *DescWRect) Push(x interface{}) {
    *pq = append(*pq, x.(*Rect))
}

func (pq *DescWRect) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    //item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

func (pq *DescHRect) Push(x interface{}) {
    *pq = append(*pq, x.(*Rect))
}

func (pq *DescHRect) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    //item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}
