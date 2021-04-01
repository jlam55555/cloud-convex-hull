package main

import (
	"convexhull/test"
	"convexhull/utils"
	"log"
	"testing"
)

const N = int(1e1)

// TestConvexHull2D tests the 2D convex hull methods
func TestConvexHull2D(t *testing.T) {
	vs := test.GenerateVertexSet(N)

	utils.PlotVertexSet2(&vs, "../../res/test.png")

	log.Fatalln("Number of vertices: ", len(vs.Vertices))
}
