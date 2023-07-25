package bigo

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Plot() {
	// Create a new plot
	p := plot.New()

	p.X.Label.Text = "n"
	p.Y.Label.Text = "Runtime"

	n := make(plotter.XYs, 100)
	bigO := make([]plotter.XYs, len(labels))

	for i := range n {
		nVal := float64(i+1) / 10
		n[i].X = nVal

		for j := range bigO {
			bigO[j] = append(bigO[j], plotter.XY{nVal, calculateComplexity(j, nVal)})
		}
	}

	lines := make([]plot.Plotter, len(bigO))
	for i := range lines {
		var err error
		lines[i], err = plotter.NewLine(bigO[i])
		if err != nil {
			panic(err)
		}
	}

	p.Add(lines...)

	for _, label := range labels {
		p.Legend.Add(label)
	}

	if err := p.Save(10*vg.Inch, 8*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}

	fmt.Println("Plot saved to plot.png")
}

func calculateComplexity(index int, n float64) float64 {
	switch index {
	case 0:
		return 1.0
	case 1:
		return math.Log(n)
	case 2:
		return n
	case 3:
		return n * math.Log(n)
	case 4:
		return n * n
	case 5:
		return n * n * n
	case 6:
		return math.Pow(2, n)
	default:
		return 0.0
	}
}

var labels = []string{"Constant", "Logarithmic", "Linear", "Log linear", "Quadratic", "Exponential", "Cubic"}
