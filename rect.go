package stripack

import "fmt"

type Rect struct {
	X, Y int
	H, W int
	id int
}

func (r *Rect) isEmpty() bool {
	return r.H <= 0 || r.W <= 0
}

func (r *Rect) String() string {
	return fmt.Sprintf("id=%v,X=%v,Y=%v,H=%v,W=%v \n",r.id,r.X,r.Y,r.H,r.W)
}

// Returns whether inner fits for packing into outer and corresponding vertical 
// flag. Note that inner is packable when at least one of its dimensions differ 
// from corresponding dimension of outer for no more than delta.
func (inner *Rect)  Packable(outer *Rect) bool {
	return (inner.W <= outer.W) && (inner.H <= outer.H)
}