package model

type HalfEdge struct {
	Head             *Vertex3
	Next, Prev, Twin *HalfEdge
	Face             *Face
}

func NewHalfEdge(face *Face, head *Vertex3) HalfEdge {
	return HalfEdge{Face: face, Head: head}
}

func (h1 *HalfEdge) SetOpposite(h2 *HalfEdge) {
	h1.Twin = h2
	h2.Twin = h1
}

func (h1 *HalfEdge) Tail() *Vertex3 {
	if h1.Twin == nil {
		return nil
	}
	return h1.Twin.Head
}

func (h1 *HalfEdge) OppositeFace() *Face {
	if h1.Twin == nil {
		return nil
	}
	return h1.Twin.Face
}

func (h1 *HalfEdge) Length() float64 {
	return h1.Head.Pos.Dist(&h1.Tail().Pos)
}

func (h1 *HalfEdge) Length2() float64 {
	return h1.Head.Pos.Dist2(&h1.Tail().Pos)
}
