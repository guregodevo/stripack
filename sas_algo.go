package stripack


import (
	"fmt"
	"container/heap"	
)

type SasAlgo struct {
	nRects *DescHRect //narrow rects
	wRects *DescWRect //wide rects
	packedRects []*Rect
	h map[int]int
	level int
}

func (algo *SasAlgo) PrettyPrint(W int, H int) {
	strip := Array(W,H)
	fmt.Printf("Printing W:%v,H:%v \n", W, H)
	for _, rect := range algo.packedRects {
		print(rect.String())
		for w := rect.X; w < rect.X + rect.W; w++ {
			fmt.Printf("strip W:%v\n", w)
			for h := rect.Y; h < rect.Y + rect.H; h++ {
				fmt.Printf("strip w:%v,h:%v \n", w, h)
		    	strip[w][h] = rect.id 	 	
			}
		}			
	}
	for h := H -1; h >= 0; h-- {
		for w := 0; w < W; w++ {
	    	print("|")
	    	if strip[w][h] == -1 {
	    		print("_")
	    	} else {
	    		print(strip[w][h])	
	    	}
	    	print("") 	 	
		}
		println("")
	}
}

//Partition the list of rectangles L = L1 ∪ L2 such that L1 is a list
//with h(Li ) > w(Li ) for all 1 ≤ i ≤ n1 ,
//while L2 is a list with w(Lj ) ≥ h(Lj ) for all 1 ≤ j ≤ n2 
func partition(rects []*Rect) (l1 *DescHRect,l2 *DescWRect) {
	l1 = &DescHRect{}
	l2 = &DescWRect{}
	heap.Init(l1)
	heap.Init(l2)
	for i := 0; i < len(rects); i++ {
		if rects[i].H > rects[i].W {
			heap.Push(l1, rects[i])  
		} else {
			heap.Push(l2, rects[i])
		}
	}
	return
}

func (v *SasAlgo) Pack(W int, rects []*Rect) int {
	n := len(rects)

	v.nRects, v.wRects = partition(rects)
	fmt.Printf("narrow rects:%v \n", v.nRects.Len())
	fmt.Printf("wide rects:%v \n", v.wRects.Len())
	fmt.Printf("Total rect:%v\n", n)

	//v.nRects contains narrow rects
	//v.wRects contains wide rects
	//Order L1 according to non-increasing height 
	//and order L2 according to non-increasing width.
	v.h = map[int]int{0:0}
	v.level = 0

	for v.nRects.Len() > 0 || v.wRects.Len() > 0 {
		//compare h(Li ) with h(Lj ) and select the rectangle with greatest height
		var rectWide, rectNarrow *Rect

		isNarrow := v.nRects.Len() != 0
		isWide := v.wRects.Len() != 0

		if isNarrow {
			rectNarrow = heap.Pop(v.nRects).(*Rect)
		}

		if isWide {
			rectWide = heap.Pop(v.wRects).(*Rect)
		}

		if (isNarrow && !isWide) || (isNarrow && isWide && rectNarrow.H > rectWide.H ) {
			//tallest rectangle is narrow
			//h(level + 1) ← h(level) + h(Li );
			v.h[v.level+1] = v.h[v.level] + rectNarrow.H 

			fmt.Printf("Pack first Narrow rect %v | level=%v \n", rectNarrow, v.level)
			//Pack the selected rectangle on the level
			v.packToRect(rectNarrow, 0, v.h[v.level])
			frame := new(Rect)
			frame.X = rectNarrow.W
			frame.Y = v.h[v.level]
			frame.H = rectNarrow.H
			frame.W = W - rectNarrow.W

			if isWide {
				heap.Push(v.wRects, rectWide) //restore rectWide
			}
			print("Frame:")
			print(frame.String())
			v.packWideRects(frame)
		} else {
			//tallest rectangle is wide
			//heap.Push(lj, rectj) //restore recti
			//h(level + 1) ← h(level) + h(Lj );
			v.h[v.level+1] = v.h[v.level] + rectWide.H

			//Pack the selected rectangle on the level
			fmt.Printf("Pack first Wide rect %v | level=%v \n", rectWide, v.level)
			v.packToRect(rectWide, 0, v.h[v.level])
			frame := new(Rect)
			frame.X = rectWide.W
			frame.Y = v.h[v.level]
			frame.H = rectWide.H
			frame.W = W - rectWide.W

			if isNarrow {
				heap.Push(v.nRects, rectNarrow) //restore rectNarrow
			}
			v.packNarrowRects(frame)
		}
		v.level = v.level + 1
	}
	return v.level
}

func (v *SasAlgo) packToRect(r *Rect, X int, Y int) {
	r.X = X
	r.Y = Y
	fmt.Printf("Packed")
	print(r.String())
	v.packedRects = append(v.packedRects, r)
}

func (v *SasAlgo) packWideRects(outer *Rect) {
	println("Pack wide rect")
	leftspace := new(Rect)
	leftspace.Y = outer.Y
	leftspace.W = outer.W //to confirm. Condition for bottom-most wide
	leftspace.X = outer.X
	leftspace.H = outer.H

	var lastPackedRect *Rect
	lastPackedRect = nil
	for v.wRects.Len() > 0 {
		wRect := heap.Pop(v.wRects).(*Rect)
		if wRect.Packable(leftspace) {
			v.packToRect(wRect, leftspace.X, leftspace.Y)
			//if rectangles of unequal widths are stacked
			if lastPackedRect != nil && lastPackedRect.W != wRect.W {
				//region R is created 
				//and narrow rectangles are packed in this region
				print("**Actual Rect")
				print(wRect.String())
				print("**last Rect")
				print(lastPackedRect.String())				
				regionR := new(Rect)
				regionR.X = wRect.X + wRect.W
				regionR.Y = wRect.Y
				regionR.W = lastPackedRect.W - wRect.W //To confirm 
				regionR.H = v.h[v.level+1] - regionR.Y //To confirm 
				print("**RegionR:")
				print(regionR.String())
				v.packNarrowRects(regionR)
			}
			lastPackedRect = wRect
			leftspace.Y = leftspace.Y + wRect.H
		} else {
			defer heap.Push(v.wRects, wRect)
		}
	}
	return		
}

func (v *SasAlgo) packNarrowRects(outer *Rect) {
	println("Pack narrow rect")
	if outer.isEmpty() {
		return
	}

	for v.nRects.Len() > 0 {
		firstNRect := heap.Pop(v.nRects).(*Rect)
		if firstNRect.Packable(outer) {
			//pack first narrow rectangle that fits height-wise and width-wise
			v.packToRect(firstNRect, outer.X, outer.Y)
			lastRect := firstNRect.clone()			
			for v.nRects.Len() > 0 {
				nRect := heap.Pop(v.nRects).(*Rect)

				leftspace := new(Rect)
				leftspace.Y = lastRect.Y + lastRect.H
				leftspace.W = lastRect.W //to confirm. Condition for bottom-most wide
				leftspace.X = lastRect.X
				leftspace.H = v.h[v.level+1] - (lastRect.H + lastRect.Y) 
			
				//search L1 for rectangle whose width is at most the width of the bottom-most narrow
				if nRect.Packable(leftspace) {
					v.packToRect(nRect, leftspace.X, leftspace.Y)
					leftspace.Y = lastRect.Y + lastRect.H
					lastRect = nRect
					//TODO shift to the right if there is no more space
					if leftspace.Y >= v.h[v.level] {
						leftspace.X = leftspace.X + nRect.W
						leftspace.Y = outer.Y
						leftspace.H = v.h[v.level] - outer.Y
						leftspace.W = outer.X + outer.W - leftspace.X   
						v.packNarrowRects(leftspace)
						break
					} 
				} else {
					defer heap.Push(v.nRects, nRect)
				}
			}
			break
		} else {
			defer heap.Push(v.nRects, firstNRect)
		}
	}
	return
}
