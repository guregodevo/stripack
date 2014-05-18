package stripack

import (
)

type OnlineAlgorithm interface {
	
	//Packing a list of rectangles into a strip of fixed width and infinite height. The
	//list of rectangles is fully specified in advance, before packing commences
	//Input : The rectangles to be packed n and the strip width W .
	//Ouput : The height H of the packing obtained in the strip.
	Pack(H int, W int, rects []*Rect, rect *Rect) (bool, *Rect) 
}