package utils

import (
	"convexhull/model"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
)

// PlotVertexSet2 plots a 2D vertex set using gonum/plot
func PlotVertexSet2(vs *model.VertexSet2, filename string) {
	plt := plot.New()
	scatter, err := plotter.NewScatter(vs)
	if err != nil {
		log.Fatalln(err)
	}

	plt.Add(scatter)

	err = plt.Save(4*vg.Inch, 4*vg.Inch, filename)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Done")
}
