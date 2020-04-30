package goarch

import (
	"github.com/melaurent/goarch/mean"
	"github.com/melaurent/goarch/volatility"
	"math"
	"testing"
)

func TestModel_Progress(t *testing.T) {
	m := NewModel(mean.NewConstantMeanProcess(10), volatility.NewGARCH(
		2.,
		0.0000,
		[]float64{0.1000},
		[]float64{0.0100},
		[]float64{0.8750}))

	samples := m.Sample(10)
	if len(samples) != 10 {
		t.Fatalf("was expecting 10 samples, got %d", len(samples))
	}
	for _, s := range samples {

		if math.Abs(s-10.) > 0.0000001 {
			t.Fatalf("was expecting 10, got %f", s)
		}
	}
}
