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
	f.EdgeCount = 3

	f.calcCentroid()
	f.calcNormal()

	return f
}

func (f *Face) calcCentroid() {
	c := NewZeroVector3()

	for h, i := f.EdgeHead.Next, 0; i < f.EdgeCount; h, i = h.Next, i+1 {
		c = c.Add(&h.Head.Pos)
	}
	f.Centroid = c.Scale(1 / float64(f.EdgeCount))
}

func (f *Face) calcNormal() {
	// for stability purposes, calculate normal from adjacent edge pairs;
	// use Newell's method, which is a little less computationally expensive
	// than directly calculating product of adjacent pairs
	// see: https://www.khronos.org/opengl/wiki/Calculating_a_Surface_Normal#Newell.27s_Method
	// the Java implementation also has an additional step to improve
	// robustness by removing any component of the normal parallel to the
	// longest edge;
	// see also: https://stackoverflow.com/a/22838372/2397327

	n, h := NewZeroVector3(), f.EdgeHead
	d2 := h.Head.Pos.Minus(&h.Prev.Head.Pos)

	for i := 0; i < f.EdgeCount; h, i = h.Next, i+1 {
		d1 := d2
		d2 = h.Next.Head.Pos.Minus(&h.Head.Pos)

		n.X += (d1.Y - d2.Y) * (d1.Z + d2.Z)
		n.Y += (d1.Z - d2.Z) * (d1.X + d2.X)
		n.Z += (d1.X - d2.X) * (d1.Y + d2.Y)
	}
	f.Normal = n.Normalize()
}

func (f *Face) GetEdge(i int) *HalfEdge {
	h := f.EdgeHead

	for ; i > 0; h, i = h.Next, i-1 {
	}
	for ; i < 0; h, i = h.Prev, i+1 {
	}

	return h
}
