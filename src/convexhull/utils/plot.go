package utils

import (
	"convexhull/model"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
)

// PlotVertex2Set plots a 2D vertex set using gonum/plot
func PlotVertex2Set(vs *model.Vertex2Set, filename string) {
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
