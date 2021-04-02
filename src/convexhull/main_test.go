package main

import (
	"convexhull/algo2d/quickhull"
	"convexhull/test"
	"convexhull/utils"
	"log"
	"testing"
)

const N = int(1e1)

// TestConvexHull2D tests the 2D convex hull methods
func TestConvexHull2D(t *testing.T) {
	vs := test.GenerateVertexSet(N)

	utils.PlotVertex2Set(vs, "../../res/test.png")

	quickhull.QuickHull2D(vs)

	log.Fatalln("Number of vertices: ", len(vs.Vertices))
}
