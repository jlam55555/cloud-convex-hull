package model

type Face struct {
	EdgeHead     *HalfEdge
	Normal       *Vector3
	ConflictList []*Vector3
}
