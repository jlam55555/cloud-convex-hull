package main

import (
	"convexhull/algo2d/quickhull2d"
	"convexhull/test"
	"convexhull/utils"
	"log"
	"testing"
)

const N = int(1e6)

// TestConvexHull2D tests the 2D convex hull methods
func TestConvexHull2D(t *testing.T) {
	vs := test.GenerateVertexSet(N)

	log.Println("Generating plot of original...")
	utils.PlotVertex2Set(vs, "../../res/test.png")

	log.Println("Performing QuickHull2D...")
	ch := quickhull2d.QuickHull2D(vs)

	log.Println("Generating plot of convex hull...")
	utils.PlotVertex2Set(ch, "../../res/ch.png")

	log.Fatal("Done")
}
