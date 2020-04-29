package goarch

import "testing"
import "github.com/melaurent/goarch/mean"
import "github.com/melaurent/goarch/volatility"

func TestModel_Progress(t *testing.T) {
	m := NewModel(mean.NewConstantMeanProcess(0), volatility.NewGARCH(
		2.,
		0.0200,
		[]float64{0.1000},
		[]float64{0.0100},
		[]float64{0.8750}))
	m.Progress(0.)
}
