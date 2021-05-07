package model

type Vertex3 struct {
	Prev, Next *Vertex3
	Pos        Vector3
}

func NewVertex3(x, y, z float64) Vertex3 {
	return Vertex3{Pos: Vector3{x, y, z}}
}

func NewVertex3FromSlice(pos [3]float64) Vertex3 {
	return Vertex3{Pos: NewVector3FromSlice(pos)}
}

func NewVertex3SliceFromSlice(poss [][3]float64) []Vertex3 {
	vs := make([]Vertex3, len(poss))
	for i, pos := range poss {
		vs[i] = NewVertex3FromSlice(pos)
	}
	return vs
}
