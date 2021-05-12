package objio

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Parse reads an obj file into a list of vertices; this is specific to this
// convex hull operation, since all face information can be discarded
func Parse(reader io.Reader) ([][3]float64, error) {
	var point [3]float64
	var points [][3]float64

	s := bufio.NewScanner(reader)

	for s.Scan() {
		ss := bufio.NewScanner(strings.NewReader(s.Text()))
		ss.Split(bufio.ScanWords)

		// ignore anything except vertex
		if word := ss.Scan(); !word || ss.Text() != "v" {
			continue
		}

		// read in xyz coordinates
		for i := 0; i < 3; i++ {
			if word := ss.Scan(); !word {
				return nil, errors.New("invalid vertex: " +
					"< 3 coordinates")
			}

			var err error
			if point[i], err = strconv.ParseFloat(
				ss.Text(), 64); err != nil {
				return nil, err
			}
		}
		points = append(points, point)
	}

	return points, nil
}

func Dump(writer io.Writer, vertices [][3]float64, faces [][3]int) error {
	w := bufio.NewWriter(writer)
	defer w.Flush()

	if _, err := w.WriteString(
		"# created by convex hull (Jonathan Lam)\n"); err != nil {
		return err
	}

	for _, v := range vertices {
		if _, err := w.WriteString(fmt.Sprintf(
			"v %f %f %f\n", v[0], v[1], v[2])); err != nil {
			return err
		}
	}

	for _, f := range faces {
		if _, err := w.WriteString(fmt.Sprintf(
			"f %d %d %d\n", f[0], f[1], f[2])); err != nil {
			return err
		}
	}

	return nil
}
