package main

import (
	"bufio"
	"fmt"
	"objio"
	"os"
	"qh3d"
)

func main() {
	file, err := os.OpenFile("res/teapot.obj", os.O_RDONLY, 0)
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

	ch, err := qh3d.QuickHull3DFromSlice(points)
	if err != nil {
		panic(err)
	}

	// convert face indices to vertex indices
	vertexMap := make(map[int]int)
	vertices := make([][3]float64, 0)
	faces := make([][3]int, 0)
	for i, v := range ch.VerticesF {
		if v {
			vertices = append(vertices,
				ch.Vertices[i].Pos.ToSlice())
			vertexMap[i] = len(vertices)
		}
	}
	for _, f := range ch.Faces {
		// currently only supporting a simplex (triangular faces)
		faces = append(faces, [3]int{
			vertexMap[f.GetEdge(0).Head.Index],
			vertexMap[f.GetEdge(1).Head.Index],
			vertexMap[f.GetEdge(2).Head.Index],
		})
	}

	file, err = os.OpenFile("res/test.obj", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err = objio.Dump(file, vertices, faces); err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}
