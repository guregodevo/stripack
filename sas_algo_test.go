package stripack

import (
	"testing"
)


func Test_SAS_Offline_Pack(t *testing.T) {

	rects := []*Rect{&Rect{id:5, H:5, W:1}, 
	&Rect{id:2, H:1, W:2}, 
	&Rect{id:1, H:1, W:1}, 
	&Rect{id:7, H:2, W:4}, 
	&Rect{id:6, H:1, W:1}, 
	&Rect{id:8, H:1, W:2},
	&Rect{id:9, H:3, W:1},
	&Rect{id:3, H:3, W:3},
	&Rect{id:4, H:4, W:2},	
	}
	algo := new(SasAlgo)

    height, rects := algo.Pack(5, rects)
    expect(t, 9, height)
    expect(t, 9, len(rects))
}
