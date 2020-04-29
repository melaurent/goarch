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
