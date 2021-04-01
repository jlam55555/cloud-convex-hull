package model

// VertexSet2 is container for 2D vertices
type VertexSet2 struct {
	Vertices []Vertex2
}

// VertexSet3 is container for 3D vertices
type VertexSet3 struct {
	Vertices []Vertex3
}

func (vs *VertexSet2) Len() int {
	return len(vs.Vertices)
}

func (vs *VertexSet2) XY(i int) (float64, float64) {
	return vs.Vertices[i].X, vs.Vertices[i].Y
}
