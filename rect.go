package stripack

import "fmt"
import "math/rand"

type Rect struct {
	X, Y int
	H, W int
	id int
}

func (r *Rect) isEmpty() bool {
	return r.H <= 0 || r.W <= 0
}

func (r *Rect) clone() *Rect {
	p := new(Rect)
	p.id = r.id
	p.H = r.H
	p.W = r.W
	p.X = r.X
	p.Y = r.Y
	return p
} 

func (r *Rect) String() string {
	return fmt.Sprintf("id=%v,X=%v,Y=%v,H=%v,W=%v \n",r.id,r.X,r.Y,r.H,r.W)
}

// Returns whether inner fits for packing into outer and corresponding vertical 
// flag. Note that inner is packable when at least one of its dimensions differ 
// from corresponding dimension of outer for no more than delta.
func (inner *Rect)  Packable(outer *Rect) bool {
	return inner.W <= outer.W && inner.H <= outer.H
}

func PrettyPrint(packedRects []*Rect, W int, H int) {
	strip := Array(W,H)
	for _, rect := range packedRects {
		for w := rect.X; w < rect.X + rect.W; w++ {
			for h := rect.Y; h < rect.Y + rect.H; h++ {
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

func GenerateRectangles(n int) []*Rect {
	res := make([]*Rect, n)
	for y := 0; y < n; y++ {
		res[y].H = rand.Int()
		res[y].W = rand.Int()
		res[y].X = rand.Int()
		res[y].Y = rand.Int()
		res[y].id = rand.Int()
	}
	return res
}

func TotalArea(rects []*Rect) int {
	var res int = 0
	for _, r := range rects {
		res += r.W * r.H
	}
	return res
}