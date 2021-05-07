package model

type HalfEdge struct {
	Tail             *Vector3
	Next, Prev, Twin *HalfEdge
	Face             *Face
}

func (h1 *HalfEdge) SetOpposite(h2 *HalfEdge) {
	h1.Twin = h2
	h2.Twin = h1
}

func (h1 *HalfEdge) Head() *Vector3 {
	if h1.Twin == nil {
		return nil
	}
	return h1.Twin.Tail
}

func (h1 *HalfEdge) OppositeFace() *Face {
	if h1.Twin == nil {
		return nil
	}
	return h1.Twin.Face
}

func (h1 *HalfEdge) Length() float64 {
	return h1.Tail.Dist(h1.Head())
}

func (h1 *HalfEdge) Length2() float64 {
	return h1.Tail.Dist2(h1.Head())
}
