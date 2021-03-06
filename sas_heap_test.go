package stripack


import (
	"container/heap"
	"testing"	
)

func Test_Narrow_Rect_Heap(t *testing.T) {
	square := &Rect{ Id:1, H:1, W:2}
	rect := &Rect{ Id:2, H:2, W:2}
	vrect := &Rect{ Id:3, H:1, W:2}

	l1 := &NarrowRectHeap{}
	heap.Init(l1)
	heap.Push(l1,rect)
	heap.Push(l1,square)
	heap.Push(l1,vrect)
	
	expect(t, 3, l1.Len())
	s := heap.Pop(l1)
	expect(t, 2, l1.Len())
	expect(t, rect, s)
}

func Test_Wide_Rect_Heap(t *testing.T) {
	square := &Rect{ H:1, W:2}
	rect := &Rect{ H:1, W:5}
	vrect := &Rect{ H:1, W:2}

	l1 := &WideRectHeap{}
	heap.Init(l1)
	heap.Push(l1,rect)
	heap.Push(l1,square)
	heap.Push(l1,vrect)
	
	expect(t, 3, l1.Len())
	s := heap.Pop(l1)
	expect(t, 2, l1.Len())
	expect(t, rect, s)
}
