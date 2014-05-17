package stripack

import (
	"testing"
)


func Test_Pack(t *testing.T) {

	rects := []*Rect{&Rect{id:5, H:5, W:1}, 
	&Rect{id:2, H:1, W:2}, 
	&Rect{id:1, H:1, W:1}, 
	&Rect{id:7, H:2, W:4}, 
	&Rect{id:6, H:1, W:1}, 
	&Rect{id:8, H:2, W:1}}
	algo := new(SasAlgo)

    algo.Pack(5, rects)
	algo.PrettyPrint(5,9)
}
