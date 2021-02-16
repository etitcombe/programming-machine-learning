package stat

import (
	"math"

	"gonum.org/v1/gonum/stat"
)

// LinearRegression wraps the function provided by gonum/stat.
func LinearRegression(xs, ys []float64) (float64, float64) {
	return stat.LinearRegression(xs, ys, nil, false)
}

// Loss returns the mean squared error for the given data points and slope w.
func Loss(xs, ys []float64, w, b float64) float64 {
	var sum float64
	for i := 0; i < len(xs); i++ {
		sum += math.Pow(Predict(xs[i], w, b)-ys[i], 2)
	}
	return sum / float64(len(xs))
}

// Predict returns the predicted value for the given data x and slope w.
func Predict(x, w, b float64) float64 {
	return x*w + b
}
