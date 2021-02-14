package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	points, err := getData("input/pizza.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = generateScatter(points)
	if err != nil {
		log.Fatal(err)
	}
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

func generateScatter(points plotter.XYs) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.X.Label.Text = "Reservations"
	p.Y.Label.Text = "Pizzas"

	err = plotutil.AddScatters(p, points)
	if err != nil {
		return err
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "output/scatter.png"); err != nil {
		return err
	}
	return nil
}
