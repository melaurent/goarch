package mean

type ARProcess struct {
	constant float64
	lags     []int
	coefs    []float64
	vals     []float64
	valIdx   int
}

func NewARProcess(constant float64, lags []int, coefs []float64) *ARProcess {
	if len(lags) != len(coefs) {
		panic("lags and coefficients need to have the same length")
	}
	maxLag := 0
	for _, l := range lags {
		if l == 0 {
			panic("invalid lag of 0")
		}
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
		// Should be m.valIdx-1 - (l - 1)
		y += m.vals[(((m.valIdx-l)%N)+N)%N] * m.coefs[j]
	}
	m.vals[m.valIdx] = y
	m.valIdx = (m.valIdx + 1) % N

	return x - y
}

func (m *ARProcess) Sample(nIntervals int) []float64 {

	samples := make([]float64, nIntervals, nIntervals)

	N := len(m.vals)
	vals := make([]float64, N, N)
	for i := 0; i < len(m.vals); i++ {
		vals[i] = m.vals[(m.valIdx+i)%N]
	}
	valIdx := 0

	for i := 0; i < nIntervals; i++ {
		y := m.constant
		for j, l := range m.lags {
			// Should be m.valIdx-1 - (l - 1)
			y += vals[(((valIdx-l)%N)+N)%N] * m.coefs[j]
		}

		vals[valIdx] = y
		valIdx = (valIdx + 1) % N
		samples[i] = y
	}

	return samples
}
