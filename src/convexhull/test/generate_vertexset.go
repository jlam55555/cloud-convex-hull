package test

import (
	"convexhull/model"
	"math/rand"
)

// GenerateVertexSet generates a random normally-distributed pointset
func GenerateVertexSet(N int) model.VertexSet2 {
	vs := model.VertexSet2{Vertices: make([]model.Vertex2, N)}

	for i := 0; i < N; i++ {
		x := rand.NormFloat64()
		y := rand.NormFloat64()

		vs.Vertices[i] = model.Vertex2{X: x, Y: y}
	}

	return vs
}
