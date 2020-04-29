package mean

type MeanProcess interface {
	Progress(val float64) float64
	Sample(nIntervals int) []float64
}
