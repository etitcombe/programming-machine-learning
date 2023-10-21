package stat

import (
	"math"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

// LinearRegression wraps the function provided by gonum/stat.
func LinearRegression(xs, ys []float64) (float64, float64) {
	return stat.LinearRegression(xs, ys, nil, false)
}

// Gradient returns the gradient for the points at w and b.
func Gradient(xs, ys []float64, w, b float64) (float64, float64) {
	var sumW float64
	var sumB float64
	for i := 0; i < len(xs); i++ {
		p := Predict(xs[i], w, b)
		sumW += xs[i] * (p - ys[i])
		sumB += p - ys[i]
	}
	return 2 * sumW / float64(len(xs)), 2 * sumB / float64(len(xs))
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

// LostMulti returns the mean squared error for the given multi-dimensional data and slope w.
func LossMulti(mx, my, w *mat.Dense) *mat.Dense {
	mx.
}

// PredictMulti returns the predicted value for the given multi-dimensional data x and slope w.
func PredictMulti(x, w *mat.Dense) *mat.Dense {
	x.Mul(x, w)
	return x
}
