package goarch

import (
	"github.com/melaurent/goarch/mean"
	"github.com/melaurent/goarch/volatility"
)

type Model struct {
	mean mean.MeanProcess
	vol  volatility.VolatilityProcess
}

func NewModel(m mean.MeanProcess, vol volatility.VolatilityProcess) *Model {
	return &Model{
		mean: m,
		vol:  vol,
	}
}

func (m *Model) Progress(val float64) {
	m.vol.Progress(m.mean.Progress(val))
}

func (m *Model) Sample(nIntervals int) []float64 {
	ys := m.mean.Sample(nIntervals)
	errors, _ := m.vol.Sample(nIntervals)

	for i := 0; i < nIntervals; i++ {
		ys[i] += errors[i]
	}

	return ys
}
