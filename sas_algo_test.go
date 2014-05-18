package stripack

import (
	"testing"
)


func Test_SAS_Offline_Pack(t *testing.T) {

	rects := []*Rect{&Rect{Id:5, H:5, W:1}, 
	&Rect{Id:2, H:1, W:2}, 
	&Rect{Id:1, H:1, W:1}, 
	&Rect{Id:7, H:2, W:4}, 
	&Rect{Id:6, H:1, W:1}, 
	&Rect{Id:8, H:1, W:2},
	&Rect{Id:9, H:3, W:1},
	&Rect{Id:3, H:3, W:3},
	&Rect{Id:4, H:4, W:2},	
	}
	algo := new(SasAlgo)

    height, rects := algo.Pack(5, rects)
    expect(t, 9, height)
    expect(t, 9, len(rects))
}
