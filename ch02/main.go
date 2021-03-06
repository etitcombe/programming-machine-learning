package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/etitcombe/programming-machine-learning/data"
	"github.com/etitcombe/programming-machine-learning/stat"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// points, err := getData("input/pizza.txt")
	// points, err := getSeparatedData("input/pizza.txt")
	xs, ys, err := getXYs("input/pizza.txt")
	if err != nil {
		log.Fatal(err)
	}

	w, b, err := train(xs, ys, 100000, 0.001)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("w = %.3f; b = %.3f\n", w, b)

	p := stat.Predict(20, w, b)
	fmt.Printf("p = %.2f\n", p)

	err = generateScatter(xs, ys, w, b, "output/scatter.png")
	if err != nil {
		log.Fatal(err)
	}

	b, w = stat.LinearRegression(xs, ys)
	fmt.Printf("w = %.3f; b = %.3f\n", w, b)

	p = stat.Predict(20, w, b)
	fmt.Printf("p = %.2f\n", p)

	err = generateScatter(xs, ys, w, b, "output/scatter2.png")
	if err != nil {
		log.Fatal(err)
	}
}

func getPlotterData(xs, ys []float64) (plotter.XYs, error) {
	if len(xs) != len(ys) {
		return nil, fmt.Errorf("the number of x values must match the number of y values")
	}

	data := make(plotter.XYs, len(xs))
	for i := 0; i < len(xs); i++ {
		data[i] = plotter.XY{X: xs[i], Y: ys[i]}
	}

	return data, nil
}

func getXYs(fileName string) ([]float64, []float64, error) {
	records, err := data.ReadSeparatedValues(fileName, ' ', 2)
	if err != nil {
		return nil, nil, err
	}

	// skip the first record because it contains the column headings
	records = records[1:]

	xs := make([]float64, len(records))
	ys := make([]float64, len(records))
	for i := 0; i < len(records); i++ {
		x, err := strconv.ParseFloat(records[i][0], 64)
		if err != nil {
			return nil, nil, err
		}
		xs[i] = x
		y, err := strconv.ParseFloat(records[i][1], 64)
		if err != nil {
			return nil, nil, err
		}
		ys[i] = y
	}

	return xs, ys, nil
}

func generateScatter(xs, ys []float64, w, b float64, fileName string) error {
	points, err := getPlotterData(xs, ys)
	if err != nil {
		log.Fatal(err)
	}

	p := plot.New()
	p.X.Min = 0
	p.X.Max = 30
	p.Y.Min = 0
	p.Y.Max = 60
	p.X.Label.Text = "Reservations"
	p.Y.Label.Text = "Pizzas"

	// fmt.Println(points)

	f := plotter.NewFunction(func(x float64) float64 { return w*x + b })
	f.Color = color.RGBA{B: 128, A: 255}

	s, err := plotter.NewScatter(points)
	if err != nil {
		return err
	}
	s.Color = color.RGBA{R: 255, B: 255, A: 255}

	p.Add(f, s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, fileName); err != nil {
		return err
	}
	return nil
}

func train(xs, ys []float64, iterations int, lr float64) (float64, float64, error) {
	if len(xs) != len(ys) {
		return 0, 0, fmt.Errorf("the number of x values must match the number of y values")
	}

	var w float64
	var b float64

	for i := 0; i < iterations; i++ {
		currentLoss := stat.Loss(xs, ys, w, b)
		if stat.Loss(xs, ys, w+lr, b) < currentLoss {
			w += lr
		} else if stat.Loss(xs, ys, w-lr, b) < currentLoss {
			w -= lr
		} else if stat.Loss(xs, ys, w, b+lr) < currentLoss {
			b += lr
		} else if stat.Loss(xs, ys, w, b-lr) < currentLoss {
			b -= lr
		} else {
			return w, b, nil
		}
	}

	return 0, 0, fmt.Errorf("couldn't converge within %d iterations", iterations)
}

/*
func getSeparatedData(fileName string) (plotter.XYs, error) {
	records, err := data.ReadSeparatedValues(fileName, ' ', 2)
	if err != nil {
		return nil, err
	}

	data := make(plotter.XYs, len(records)-1) // reduce by 1 because we ignore the first record

	for i := 1; i < len(records); i++ { // skip the first record because it contains the column headings
		v1, err := strconv.ParseFloat(records[i][0], 64)
		if err != nil {
			return nil, err
		}
		v2, err := strconv.ParseFloat(records[i][1], 64)
		if err != nil {
			return nil, err
		}
		data[i-1] = plotter.XY{X: v1, Y: v2}
	}

	return data, nil
}

func getData(fileName string) (plotter.XYs, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make(plotter.XYs, 0, 30)

	// We use a regex so that we can skip over blank lines or the header line
	re, err := regexp.Compile(`^(\d+)\s+(\d+)$`)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}
		v1, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			return nil, err
		}
		v2, err := strconv.ParseFloat(matches[2], 64)
		if err != nil {
			return nil, err
		}
		data = append(data, plotter.XY{X: v1, Y: v2})
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
*/
