package model

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

func NewZeroVector3() Vector3 {
	return Vector3{0, 0, 0}
}

func NewVector3FromSlice(v [3]float64) Vector3 {
	return Vector3{v[0], v[1], v[2]}
}

func NewVector3SliceFromSlice(poss [][3]float64) []Vector3 {
	vs := make([]Vector3, len(poss))
	for i, pos := range poss {
		vs[i] = NewVector3FromSlice(pos)
	}
	return vs
}

func (v1 *Vector3) Get(i int) float64 {
	if i == 0 {
		return v1.X
	} else if i == 1 {
		return v1.Y
	} else {
		return v1.Z
	}
}

func (v1 *Vector3) Add(v2 *Vector3) Vector3 {
	return Vector3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 *Vector3) Minus(v2 *Vector3) Vector3 {
	return Vector3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v1 *Vector3) Scale(sf float64) Vector3 {
	return Vector3{sf * v1.X, sf * v1.Y, sf * v1.Z}
}

func (v1 *Vector3) Norm() float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
}

func (v1 *Vector3) Normalize() Vector3 {
	return v1.Scale(1 / v1.Norm())
}

func (v1 *Vector3) Norm2() float64 {
	return v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z
}

func (v1 *Vector3) Dist(v2 *Vector3) float64 {
	diff := v1.Minus(v2)
	return diff.Norm()
}

func (v1 *Vector3) Dist2(v2 *Vector3) float64 {
	diff := v1.Minus(v2)
	return diff.Norm2()
}

func (v1 *Vector3) Dot(v2 *Vector3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 *Vector3) Cross(v2 *Vector3) Vector3 {
	return Vector3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 *Vector3) ToString() string {
	return fmt.Sprintf("(%f, %f, %f)", v1.X, v1.Y, v1.Z)
}
