package main

import (
	"fmt"
	"qh3d"
)

func main() {
	points := [][3]float64{
		{0, 0, 0},
		{0, 1, 1},
		{0, 1, 0},
		{1, 0, 0},
		{1, 0, 1},
		{0, 0, 1},
		{1, 1, 1},
		{1, 1, 0},
		{0.5, 0.5, 0.5},
	}

	ch, err := qh3d.QuickHull3DFromSlice(points)
	if err != nil {
		panic(err)
	}

	fmt.Println(ch)
	fmt.Println("Done.")
}
