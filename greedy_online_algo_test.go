package stripack

import (
	"testing"
)

func Test_Greedy_Online_Pack(t *testing.T) {

	rects := []*Rect{&Rect{Id:2, H:1, W:2, X:0, Y:0}, 
	&Rect{Id:3, H:3, W:3, X:2, Y:0},
	&Rect{Id:4, H:1, W:2, X:0, Y:1},	
	}
	algo := new(GreedyOnlineAlgo)


	rect := &Rect{Id:5, H:1, W:1}
    isPacked, packedRect := algo.Pack(5,5, rects, rect)
    
    expect(t, true, isPacked)
    expect(t, 0, packedRect.X)
    expect(t, 2, packedRect.Y)
    PrettyPrint(append(rects, packedRect), 5,5)
}

func Test_Greedy_Online_Pack_Too_Much_Rect(t *testing.T) {
	rects := []*Rect{&Rect{Id:1, H:2, W:3, X:0, Y:0}, 
	&Rect{Id:2, H:1, W:1, X:3, Y:1},
	&Rect{Id:3, H:1, W:1, X:4, Y:0},
	&Rect{Id:4, H:1, W:1, X:4, Y:1},
	}
	algo := new(GreedyOnlineAlgo)

	PrettyPrint(rects,5, 10)
	rect := &Rect{Id:11, H:1, W:1}
    isPacked, packedRect := algo.Pack(10,5, rects, rect)
    
    expect(t, true, isPacked)
    expect(t, 0, packedRect.X)
    expect(t, 2, packedRect.Y)
    PrettyPrint(append(rects, packedRect), 5, 10)
}
