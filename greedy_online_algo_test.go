package stripack

import (
	"testing"
)


func Test_Greedy_Online_Pack(t *testing.T) {

	rects := []*Rect{&Rect{id:2, H:1, W:2, X:0, Y:0}, 
	&Rect{id:3, H:3, W:3, X:2, Y:0},
	&Rect{id:4, H:1, W:2, X:0, Y:1},	
	}
	algo := new(GreedyOnlineAlgo)


	rect := &Rect{id:5, H:1, W:1}
    isPacked, packedRect := algo.Pack(5,5, rects, rect)
    
    expect(t, true, isPacked)
    expect(t, 0, packedRect.X)
    expect(t, 2, packedRect.Y)
    PrettyPrint(append(rects, packedRect), 5,5)
}