package main

import (
	"bufio"
	"fmt"
	"github.com/golang/geo/r3"
	"github.com/markus-wa/quickhull-go"
	"objio"
	"os"
)

// QuickHullGoTest tries out the prewritten Go library for quickhull
func QuickHullGoTest(points [][3]float64) ([][3]float64, [][3]int) {
	// convert points to r3.Vector
	pointsR3 := make([]r3.Vector, len(points))
	for i, v := range points {
		pointsR3[i] = r3.Vector{v[0], v[1], v[2]}
	}

	qh := quickhull.QuickHull{}

	ch := qh.ConvexHull(pointsR3, true, false, 0)

	// convert to vertices, faces as needed for objio
	vertices := make([][3]float64, 0)
	faces := make([][3]int, 0)
	for i, triangle := range ch.Triangles() {
		vertices = append(vertices,
			[3]float64{triangle[0].X, triangle[0].Y, triangle[0].Z},
			[3]float64{triangle[1].X, triangle[1].Y, triangle[1].Z},
			[3]float64{triangle[2].X, triangle[2].Y, triangle[2].Z},
		)
		faces = append(faces, [3]int{3*i + 1, 3*i + 2, 3*i + 3})
	}

	return vertices, faces
}

func main() {
	file, err := os.OpenFile("res/cow-nonormals.obj", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	points, err := objio.Parse(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}

	//points := [][3]float64{
	//	{0, 0, 0},
	//	{0, 1, 1},
	//	{0, 1, 0},
	//	{1, 0, 0},
	//	{1, 0, 1},
	//	{0, 0, 1},
	//	{1, 1, 1},
	//	{1, 1, 0},
	//	{0.5, 0.5, 0.5},
	//}

	vertices, faces := QuickHullGoTest(points)

	//ch, err := qh3d.QuickHull3DFromSlice(points)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// convert face indices to vertex indices
	//vertexMap := make(map[int]int)
	//vertices := make([][3]float64, 0)
	//faces := make([][3]int, 0)
	//for i, v := range ch.VerticesF {
	//	if v {
	//		vertices = append(vertices,
	//			ch.Vertices[i].Pos.ToSlice())
	//		vertexMap[i] = len(vertices)
	//	}
	//}
	//for _, f := range ch.Faces {
	//	// currently only supporting a simplex (triangular faces)
	//	faces = append(faces, [3]int{
	//		vertexMap[f.GetEdge(0).Head.Index],
	//		vertexMap[f.GetEdge(1).Head.Index],
	//		vertexMap[f.GetEdge(2).Head.Index],
	//	})
	//}

	file, err = os.OpenFile("res/test.obj", os.O_WRONLY|os.O_CREATE|
		os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err = objio.Dump(file, vertices, faces); err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}
