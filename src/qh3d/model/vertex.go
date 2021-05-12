package model

type Vertex3 struct {
	Index int
	Pos   Vector3
}

func NewFromVector3(index int, v Vector3) Vertex3 {
	return Vertex3{index, v}
}

func NewFromVector3Slice(vs []Vector3) []Vertex3 {
	vsNew := make([]Vertex3, len(vs))
	for i, v := range vs {
		vsNew[i] = NewFromVector3(i, v)
	}
	return vsNew
}
