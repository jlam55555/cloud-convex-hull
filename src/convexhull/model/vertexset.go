package model

// Vertex2Set is container for 2D vertices
type Vertex2Set struct {
	Vertices []Vertex2
}

// Vertex3Set is container for 3D vertices
type Vertex3Set struct {
	Vertices []Vertex3
}

// Vertex2Subset contains a subset of the vertices of a vertex set for
// recursive processing
type Vertex2Subset []*Vertex2

// Len returns the length for the gonum/plot library
func (vs *Vertex2Set) Len() int {
	return len(vs.Vertices)
}

// XY returns the nth coordinate for the gonum/plot library
func (vs *Vertex2Set) XY(i int) (float64, float64) {
	return vs.Vertices[i].X, vs.Vertices[i].Y
}

// ToVertexSubset creates a new Vertex2Subset indicating all the elements of the
// underlying Vertex2Set
func (vs *Vertex2Set) ToVertexSubset() Vertex2Subset {
	var is Vertex2Subset = make([]*Vertex2, vs.Len())

	// fill in all pointers
	for i, _ := range vs.Vertices {
		is[i] = &vs.Vertices[i]
	}

	return is
}

// Append is a helper function to shorten doing vss = append(vss, ...)
// and to keep Vertex2Subset general in case it becomes a struct
func (vss *Vertex2Subset) Append(v *Vertex2) {
	*vss = append(*vss, v)
}

// ToVertexSet converts a Vertex2Subset to a Vertex2Set; copies over all
// elements from the underlying Vertex2Set
func (vss Vertex2Subset) ToVertexSet() *Vertex2Set {
	vs := Vertex2Set{Vertices: make([]Vertex2, len(vss))}

	for i, v := range vss {
		vs.Vertices[i] = *v
	}

	return &vs
}
