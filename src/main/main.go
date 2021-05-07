package main

import (
	"fmt"
	"qh3d"
	"qh3d/model"
)

func main() {
	if err := qh3d.QuickHull3D(model.NewVertex3SliceFromSlice([][3]float64{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 1},
		{1, 1, 0},
		{0.5, 0.5, 0.5},
	})); err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}
