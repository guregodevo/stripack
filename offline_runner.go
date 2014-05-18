package stripack

import (
	"fmt"
)

type OfflineAlgorithm interface {
	
	//Packing a list of rectangles into a strip of fixed width and infinite height. The
	//list of rectangles is fully specified in advance, before packing commences
	//Input : The rectangles to be packed n and the strip width W .
	//Ouput : The height H of the packing obtained in the strip.
	Pack(W int, rects []*Rect) (int,[]*Rect) 
}

// Returns uncovered area.
func Run(n int, W int) (uncovered_area int) {
	rects := GenerateRectangles(n)
	var algo OfflineAlgorithm = new(SasAlgo)

	H,rects := algo.Pack(W, rects)
	total_area := TotalArea(rects)
	fmt.Printf("Solution height = %v\nTotal area = %0.9v\n", H, total_area)
	uncovered_area = H * W - total_area
	fmt.Printf("Uncovered area = %0.9v\n", uncovered_area)
	PrettyPrint(rects, W,H)
	return
}