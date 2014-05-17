package stripack


import (
	"container/heap"
	"testing"	
)

func Test_H_Heap(t *testing.T) {
	square := &Rect{ id:1, H:1, W:2}
	rect := &Rect{ id:2, H:2, W:2}
	vrect := &Rect{ id:3, H:1, W:2}

	l1 := &DescHRect{}
	heap.Init(l1)
	heap.Push(l1,rect)
	heap.Push(l1,square)
	heap.Push(l1,vrect)
	
	expect(t, 3, l1.Len())
	s := heap.Pop(l1)
	expect(t, 2, l1.Len())
	expect(t, rect, s)
}

func Test_W_Heap(t *testing.T) {
	square := &Rect{ H:1, W:2}
	rect := &Rect{ H:1, W:5}
	vrect := &Rect{ H:1, W:2}

	l1 := &DescWRect{}
	heap.Init(l1)
	heap.Push(l1,rect)
	heap.Push(l1,square)
	heap.Push(l1,vrect)
	
	expect(t, 3, l1.Len())
	s := heap.Pop(l1)
	expect(t, 2, l1.Len())
	expect(t, rect, s)
}
