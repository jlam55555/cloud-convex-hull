package qh3d

import (
	"errors"
	"qh3d/model"
)

const DoublePrecision = 2.2204460492503131e-16

// QuickHull3D is the entrypoint to the quickhull algorithm
func QuickHull3D(points []model.Vertex3) error {
	if len(points) < 4 {
		return errors.New("fewer than four points specified")
	}

	// TODO: remove
	// initBuffers

	// TODO: remove
	// setPoints

	// TODO: add any preprocessing steps here

	buildHull()

	return nil
}

// buildHull is the start of the true algorithm after any preprocessing steps
func buildHull() {
	// buildInitialHull()

	//nextVertex = getNextConflictVertex()
	//for nextVertex = getNextConflictVertex(); nextVertex != nil {
	//	addVertexToHull()
	//}
}
