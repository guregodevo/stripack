package stripack

import (
	"math/rand"
	"fmt"
)

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

type Algorithm interface {
	
	//Packing a list of rectangles into a strip of fixed width and infinite height. The
	//list of rectangles is fully specified in advance, before packing commences
	//Input : The rectangles to be packed n and the strip width W .
	//Ouput : The height H of the packing obtained in the strip.
	Pack(W int, rects []*Rect) int 
}

// Returns uncovered area.
func Run(n int, W int) (uncovered_area int) {
	rects := GenerateRectangles(n)
	var algo Algorithm = new(SasAlgo)

	H := algo.Pack(W, rects)
	total_area := TotalArea(rects)
	fmt.Printf("Solution height = %v\nTotal area = %0.9v\n", H, total_area)
	uncovered_area = H * W - total_area
	fmt.Printf("Uncovered area = %0.9v\n", uncovered_area)
	
	return
}