package quickhull

import (
	"convexhull/model"
	"convexhull/utils"
)

// distToLine calculates a measure of distance (not properly normalized) between
// a line and a point
// see: https://math.stackexchange.com/a/274728/96244
// and: https://en.wikipedia.org/wiki/Distance_from_a_point_to_a_line
func distToLine(l *model.Line2, v *model.Vertex2) float64 {
	return (v.X-l.V1.X)*(l.V2.Y-l.V1.Y) - (v.Y-l.V1.Y)*(l.V2.X-l.V1.X)
}

// quickHull2DRec is the recursive part of the QuickHull algorithm
func quickHull2DRec(vss *model.Vertex2Subset, v1 *model.Vertex2,
	v2 *model.Vertex2) []int {

	// find point furthest from the oriented line

	return make([]int, 0)
}

// partitionPoints partitions a VertexSubset by the dividing line
// through v1 and v2
func partitionPoints(vss *model.Vertex2Subset, v1, v2 *model.Vertex2) (
	model.Vertex2Subset, model.Vertex2Subset) {

	vss1, vss2 := model.Vertex2Subset{}, model.Vertex2Subset{}
	div := model.Line2{V1: v1, V2: v2}

	for _, v := range *vss {
		// don't include endpoints of line in partition
		if v == v1 || v == v2 {
			continue
		}

		// distToLine is signed perpendicular distance: we only care
		// about the sign here
		if distToLine(&div, v) > 0 {
			vss1.Append(v)
		} else {
			vss2.Append(v)
		}
	}

	return vss1, vss2
}

// QuickHull2D calculates the 2D convex hull
func QuickHull2D(vs *model.Vertex2Set) model.Vertex2Set {

	// find point furthest to the left and right
	xMin, xMax := &vs.Vertices[0], &vs.Vertices[0]
	for _, v := range vs.Vertices {
		if v.X < xMin.X {
			xMin = &v
		}
		if v.X > xMax.X {
			xMax = &v
		}
	}

	// convert input vertex set to Vertex2Subset
	vss := vs.ToVertexSubset()

	// partition points
	vss1, vss2 := partitionPoints(vss, xMin, xMax)

	// call recursive method
	utils.PlotVertex2Set(vss1.ToVertexSet(), "../../res/vss1.png")
	utils.PlotVertex2Set(vss2.ToVertexSet(), "../../res/vss2.png")

	return model.Vertex2Set{}
}
