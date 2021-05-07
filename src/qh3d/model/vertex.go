package model

type Vertex3 struct {
	Pos Vector3
}

func NewFromVector3(v Vector3) Vertex3 {
	return Vertex3{v}
}

func NewFromVector3Slice(vs []Vector3) []Vertex3 {
	vsNew := make([]Vertex3, len(vs))
	for i, v := range vs {
		vsNew[i] = NewFromVector3(v)
	}
	return vsNew
}
