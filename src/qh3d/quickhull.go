// Package qh3d includes a 3-D QuickHull implementation closely based off of
// Dirk Gregorius' Steam presentation and John Lloyd's Java implementation
//
// Gregorius: http://media.steampowered.com/apps/valve/2014/DirkGregorius_ImplementingQuickHull.pdf
// Lloyd: https://www.cs.ubc.ca/~lloyd/java/quickhull3d.html
package qh3d

import (
	"errors"
	"fmt"
	"math"
	"qh3d/model"
)

const doublePrecision = 2.2204460492503131e-16
const debug = true

// ConvexHull3 stores information about the current convex hull
type ConvexHull3 struct {
	Vertices  []model.Vertex3
	tolerance float64
}

// QuickHull3DFromSlice is a wrapper around QuickHull3D that takes a slice
// of [3]float coordinates
func QuickHull3DFromSlice(vs [][3]float64) (ConvexHull3, error) {
	return QuickHull3D(model.NewVector3SliceFromSlice(vs))
}

// QuickHull3D is the entrypoint to the quickhull algorithm
func QuickHull3D(vs []model.Vector3) (ConvexHull3, error) {
	ch := ConvexHull3{}

	if len(vs) < 4 {
		return ch, errors.New("fewer than four points specified")
	}

	ch.Vertices = model.NewFromVector3Slice(vs)

	// TODO: remove
	// setPoints

	// TODO: add any preprocessing steps here

	if err := buildHull(&ch); err != nil {
		return ch, err
	}

	return ch, nil
}

// buildHull is the start of the true algorithm after any preprocessing steps
func buildHull(ch *ConvexHull3) error {
	if err := buildInitialHull(ch); err != nil {
		return err
	}

	return nil

	//nextVertex = getNextConflictVertex()
	//for nextVertex = getNextConflictVertex(); nextVertex != nil {
	//	addVertexToHull()
	//}
}

// buildInitialHull calculates the original simplex that must be part of the
// final hull
func buildInitialHull(ch *ConvexHull3) error {
	// find min and max points in each dimension
	var min, max [3]model.Vertex3

	for i := 0; i < 3; i++ {
		min[i] = ch.Vertices[0]
		max[i] = ch.Vertices[0]
	}

	for _, v := range ch.Vertices {
		for i := 0; i < 3; i++ {
			if v.Pos.Get(i) < min[i].Pos.Get(i) {
				min[i] = v
			} else if v.Pos.Get(i) > max[i].Pos.Get(i) {
				max[i] = v
			}
		}
	}

	// calculate error tolerance; formula from original quickhull paper
	// gofmt forced this ugliness
	ch.tolerance = 3 * doublePrecision * (math.Max(
		math.Abs(max[0].Pos.X), math.Abs(min[0].Pos.X)) +
		math.Max(math.Abs(max[1].Pos.Y), math.Abs(min[1].Pos.Y)) +
		math.Max(math.Abs(max[2].Pos.Z), math.Abs(min[2].Pos.Z)))

	// calculate original simplex
	maxDist, maxDistDim := 0., 0
	for i := 0; i < 3; i++ {
		if max[i].Pos.Get(i)-min[i].Pos.Get(i) > maxDist {
			maxDist = max[i].Pos.Get(i) - min[i].Pos.Get(i)
			maxDistDim = i
		}
	}

	if maxDist <= ch.tolerance {
		return errors.New("input elements appear coincident")
	}

	var simplex [4]model.Vertex3

	// furthest points must be on the original simplex
	simplex[0] = min[maxDistDim]
	simplex[1] = max[maxDistDim]

	// find third point furthest from line l01
	l01 := simplex[0].Pos.Minus(&simplex[1].Pos)
	maxDist = 0
	for _, v := range ch.Vertices {
		diff1 := v.Pos.Minus(&simplex[0].Pos)
		xprod := diff1.Cross(&l01)
		diff2 := xprod.Norm2()
		if diff2 > maxDist {
			maxDist = diff2
			simplex[2] = v
		}
	}

	// using same tolerance as Java implementation
	if math.Sqrt(maxDist) < 100*ch.tolerance {
		return errors.New("input points appear collinear")
	}

	// calculate normal vector to plane formed by first three points of
	// simplex
	// TODO: Java implementation has an additional error correction step
	// 	that is not implemented here
	diff2 := simplex[2].Pos.Minus(&simplex[0].Pos)
	xprod := diff2.Cross(&l01)
	nrml := xprod.Normalize()
	maxDist = 0
	for _, v := range ch.Vertices {
		diff1 := v.Pos.Minus(&simplex[0].Pos)
		diff2 := diff1.Dot(&nrml)
		if diff2 > maxDist {
			maxDist = diff2
			simplex[3] = v
		}
	}

	if math.Sqrt(maxDist) < 100*ch.tolerance {
		return errors.New("input points appear coplanar")
	}

	if debug {
		fmt.Printf("original simplex:\n%s\n%s\n%s\n%s\n",
			simplex[0].Pos.ToString(), simplex[1].Pos.ToString(),
			simplex[2].Pos.ToString(), simplex[3].Pos.ToString())
	}

	// simplex vertices have been found, generate faces

	return nil
}
