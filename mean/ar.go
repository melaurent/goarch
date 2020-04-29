package mean

type ARProcess struct {
	constant float64
	lags     []int
	coefs    []float64
	vals     []float64
	valIdx   int
}

func NewARProcess(constant float64, lags []int, coefs []float64) *ARProcess {
	maxLag := 0
	for _, l := range lags {
		if l > maxLag {
			maxLag = l
		}
	}

	return &ARProcess{
		constant: constant,
		lags:     lags,
		coefs:    coefs,
		vals:     make([]float64, maxLag, maxLag),
		valIdx:   0,
	}
}

func (m *ARProcess) Progress(x float64) float64 {
	// Generate next val and return residual

	N := len(m.vals)
	y := m.constant
	for j, l := range m.lags {
		y += m.vals[(m.valIdx-1-l)%N] * m.coefs[j]
	}
	m.vals[m.valIdx] = y
	m.valIdx = (m.valIdx + 1) % N

	return x - y
}

func (m *ARProcess) Sample(nIntervals int) []float64 {

	samples := make([]float64, nIntervals, nIntervals)

	N := cap(m.vals)
	vals := make([]float64, N, N)
	for i := 0; i < len(m.vals); i++ {
		vals[N-i-1] = m.vals[(m.valIdx-i-1)%N]
	}
	valIdx := 0

	y := m.constant
	for i := 0; i < nIntervals; i++ {
		for j, l := range m.lags {
			y += vals[(valIdx-1-l)%N] * m.coefs[j]
		}

		vals[valIdx] = y
		valIdx = (valIdx + 1) % N
		samples[i] = y
	}

	return samples
}
