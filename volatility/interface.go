package volatility

type VolatilityProcess interface {
	Progress(residual float64)
	Sample(nIntervals int) []float64
}
