package mean

type EMAProcess struct {
	exp float64
	ema float64
}

func NewEMAProcess(exp float64) *EMAProcess {
	return &EMAProcess{
		exp: exp,
		ema: 0,
	}
}

func (m *EMAProcess) Progress(val float64) float64 {
	m.ema = m.exp*val + (1-m.exp)*m.ema
	return val - m.ema
}
