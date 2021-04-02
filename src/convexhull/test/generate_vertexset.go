package test

import (
	"convexhull/model"
	"math/rand"
)

// GenerateVertexSet generates a random normally-distributed pointset
func GenerateVertexSet(N int) *model.Vertex2Set {
	vs := model.Vertex2Set{Vertices: make([]model.Vertex2, N)}

	for i := 0; i < N; i++ {
		x := rand.NormFloat64()
		y := rand.NormFloat64()

		vs.Vertices[i] = model.Vertex2{X: x, Y: y}
	}

	return &vs
}
