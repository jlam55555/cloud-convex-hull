package model

type Face struct {
	EdgeHead         *HalfEdge
	ConflictList     []*Vertex3
	EdgeCount        int
	Normal, Centroid Vector3
}

func NewTriangleFace(v0, v1, v2 *Vertex3) Face {
	f := Face{}
	e0 := NewHalfEdge(&f, v0)
	e1 := NewHalfEdge(&f, v1)
	e2 := NewHalfEdge(&f, v2)

	e0.Next = &e1
	e1.Next = &e2
	e2.Next = &e0
	e0.Prev = &e2
	e1.Prev = &e0
	e2.Prev = &e1

	f.EdgeHead = &e0

	f.calcCentroid()
	f.calcNormal()

	return f
}

func (f *Face) calcCentroid() {
	c := NewZeroVector3()
	for h := f.EdgeHead.Next; h != f.EdgeHead; h = h.Next {
		c = c.Add(&h.Head.Pos)
	}
	f.Centroid = c.Scale(1 / float64(f.EdgeCount))
}

func (f *Face) calcNormal() {
	// for stability purposes, calculate normal from adjacent edge pairs
	// see: https://www.khronos.org/opengl/wiki/Calculating_a_Surface_Normal#Newell.27s_Method

	// TODO: trace out this algorithm

	//n, h1, h2 := NewZeroVector3(), f.EdgeHead.Next, f.EdgeHead.Next.Next
	//
	//for h1 != f.EdgeHead {
	//
	//}
}
