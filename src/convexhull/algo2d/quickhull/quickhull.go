package quickhull

import (
	"convexhull/model"
)

// distToLine calculates a measure of distance (not properly normalized) between
// a line and a point
// see: https://math.stackexchange.com/a/274728/96244
// and: https://en.wikipedia.org/wiki/Distance_from_a_point_to_a_line
func distToLine(l *model.Line2, v *model.Vertex2) float64 {
	return (v.X-l.V1.X)*(l.V2.Y-l.V1.Y) - (v.Y-l.V1.Y)*(l.V2.X-l.V1.X)
}

// quickHull2DRec is the recursive part of the QuickHull algorithm
func quickHull2DRec(vss model.Vertex2Subset, v1 *model.Vertex2,
	v2 *model.Vertex2) model.Vertex2Subset {

	// degenerate case: has to be contained within convex hull
	if len(vss) < 2 {
		return vss
	}

	// find point furthest from the oriented line
	div := &model.Line2{V1: v1, V2: v2}
	maxDist, vFurthest := 0., vss[0]
	for i, v := range vss {
		dist := distToLine(div, v)
		if dist > maxDist {
			maxDist = dist
			vFurthest = vss[i]
		}
	}

	// TODO: inefficient to calculate both partitions when only one is
	// 	needed; should modify partitionPoints
	vss1, _ := partitionPoints(vss, v1, vFurthest)
	vss2, _ := partitionPoints(vss, vFurthest, v2)

	ch1 := quickHull2DRec(vss1, v1, vFurthest)
	ch2 := quickHull2DRec(vss2, vFurthest, v2)

	ch := append(ch1, ch2...)
	ch = append(ch, vFurthest)

	return ch
}

// partitionPoints partitions a VertexSubset by the dividing line
// through v1 and v2
func partitionPoints(vss model.Vertex2Subset, v1, v2 *model.Vertex2) (
	model.Vertex2Subset, model.Vertex2Subset) {

	vss1, vss2 := model.Vertex2Subset{}, model.Vertex2Subset{}
	div := &model.Line2{V1: v1, V2: v2}

	for _, v := range vss {
		// don't include endpoints of line in partition
		if v == v1 || v == v2 {
			continue
		}

		// distToLine is signed perpendicular distance: we only care
		// about the sign here
		if distToLine(div, v) > 0 {
			vss1.Append(v)
		} else {
			vss2.Append(v)
		}
	}

	return vss1, vss2
}

// QuickHull2D calculates the 2D convex hull
func QuickHull2D(vs *model.Vertex2Set) *model.Vertex2Set {

	// find point furthest to the left and right
	xMin, xMax := &vs.Vertices[0], &vs.Vertices[0]
	for i, v := range vs.Vertices {
		if v.X < xMin.X {
			xMin = &vs.Vertices[i]
		}
		if v.X > xMax.X {
			xMax = &vs.Vertices[i]
		}
	}

	// convert input vertex set to Vertex2Subset
	vss := vs.ToVertexSubset()

	// partition points
	vss1, vss2 := partitionPoints(vss, xMin, xMax)

	// call recursive method
	//utils.PlotVertex2Set(&model.Vertex2Set{
	//	Vertices: []model.Vertex2{*xMin, *xMax},
	//}, "../../res/edges.png")
	//utils.PlotVertex2Set(vss1.ToVertexSet(), "../../res/vss1.png")
	//utils.PlotVertex2Set(vss2.ToVertexSet(), "../../res/vss2.png")

	ch1 := quickHull2DRec(vss1, xMin, xMax)
	ch2 := quickHull2DRec(vss2, xMax, xMin)

	ch := append(ch1, ch2...)
	ch = append(ch, xMin)
	ch = append(ch, xMax)

	return ch.ToVertexSet()
}
