package mean

import (
	"math"
	"testing"
)

func TestNewARProcess(t *testing.T) {
	N := 500

	m := NewARProcess(1., []int{1, 2}, []float64{0.5, -0.5})
	sampleMeans := m.Sample(N)

	lastValue := 0.
	lastLastValue := 0.
	means := make([]float64, N, N)
	for t := 0; t < N; t++ {
		means[t] = 1. + 0.5*lastValue - 0.5*lastLastValue
		lastLastValue = lastValue
		lastValue = means[t]
	}

	for i := 0; i < N; i++ {
		if math.Abs(means[i]-sampleMeans[i]) > 0.00001 {
			t.Fatalf("got different means at %d: %f %f", i, means[i], sampleMeans[i])
		}
	}
}

func TestARProcess_Progress(t *testing.T) {
	N := 500

	m := NewARProcess(1., []int{1, 2}, []float64{0.5, -0.5})

	lastValue := 0.
	lastLastValue := 0.
	means := make([]float64, 2*N, 2*N)
	for t := 0; t < N; t++ {
		means[t] = 1. + 0.5*lastValue - 0.5*lastLastValue
		lastLastValue = lastValue
		lastValue = means[t]
		m.Progress(means[t])
	}

	for t := N; t < 2*N; t++ {
		means[t] = 1. + 0.5*lastValue - 0.5*lastLastValue
		lastLastValue = lastValue
		lastValue = means[t]
	}
	sampleMeans := m.Sample(N)

	for i := 0; i < N; i++ {
		if math.Abs(means[N+i]-sampleMeans[i]) > 0.00001 {
			t.Fatalf("got different means at %d: %f %f", i, means[N+i], sampleMeans[i])
		}
	}
}
