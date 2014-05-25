package stripack

import (
)

type GreedyOnlineAlgo struct {
}

func (v *GreedyOnlineAlgo) Pack(H int, W int, packedRects []*Rect, rect *Rect) (bool, *Rect) {
	array := Array(W, H)
	packedRect := rect.clone()
	for _, r := range packedRects {
		for x:= r.X; x < r.X + r.W; x++ {
			if x < W {
				for y:= r.Y; y < r.Y + r.H; y++ {
					if y < H {
						array[x][y] = 1 
					}
				}
			}
		}
	}
	for x:= 0; x < W; x++ {
		for y:= 0; y < H; y++ {
			isSpace := true 
			for rectx:= x; rectx < x + rect.W; rectx++ {
				for recty:= y; recty < y + rect.H; recty++ {
					if rectx < W && recty < H && array[rectx][recty] == 1 {
						isSpace = false
						break
					}
				}
			}
			if isSpace {
				packedRect.X = x
				packedRect.Y = y
				return true, packedRect
			}
		}
	}
	return false, nil
}


