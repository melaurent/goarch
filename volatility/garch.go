package volatility

import (
	"math"
	"math/rand"
)

type GARCH struct {
	omega       float64
	alpha       []float64
	gamma       []float64
	beta        []float64
	residuals   []float64
	persistence float64
	power       float64
	resIdx      int
	errSampler  func() float64
}

func NewGARCH(power float64, omega float64, alpha, gamma, beta []float64) *GARCH {
	maxDelay := len(alpha)
	if len(gamma) > maxDelay {
		maxDelay = len(gamma)
	}

	residuals := make([]float64, maxDelay, maxDelay)

	var persistence float64 = 0
	for _, a := range alpha {
		persistence += a * 0.5
	}
	for _, g := range gamma {
		persistence += g * 0.5
	}
	for _, b := range beta {
		persistence += b
	}

	return &GARCH{
		omega:       omega,
		alpha:       alpha,
		gamma:       gamma,
		beta:        beta,
		residuals:   residuals,
		persistence: persistence,
		power:       power,
		resIdx:      0,
		errSampler:  rand.NormFloat64,
	}
}

func (m *GARCH) SetErrorSampler(errSampler func() float64) {
	m.errSampler = errSampler
}

func (m *GARCH) Progress(residual float64) {
	N := len(m.residuals)
	m.residuals[m.resIdx] = residual
	m.resIdx = (m.resIdx + 1) % N
}

func (m *GARCH) Sample(nIntervals int) ([]float64, []float64) {
	M := len(m.beta)
	fsigmas := make([]float64, M, M)
	for i := 0; i < M; i++ {
		fsigmas[i] = m.omega / (1 - m.persistence)
	}
	sigIdx := 0

	N := len(m.residuals)
	residuals := make([]float64, N, N)
	for i := 0; i < N; i++ {
		residuals[i] = m.residuals[i+m.resIdx]
	}
	resIdx := 0

	errs := make([]float64, nIntervals, nIntervals)
	sigmas2 := make([]float64, nIntervals, nIntervals)

	for i := 0; i < nIntervals; i++ {
		fsigma := m.omega
		for j := 0; j < len(m.alpha); j++ {
			fsigma += m.alpha[j] * math.Pow(residuals[(resIdx-j-1)%N], m.power)
		}

		for j := 0; j < len(m.gamma); j++ {
			if m.residuals[(resIdx-j-1)%N] < 0 {
				fsigma += m.gamma[j] * math.Pow(-residuals[(resIdx-j-1)%N], m.power)
			}
		}

		for j := 0; j < len(m.beta); j++ {
			fsigma += m.beta[j] * fsigmas[(sigIdx-j-1)%M]
		}

		fsigmas[sigIdx] = fsigma
		sigIdx = (sigIdx + 1) % M
		sigma2 := math.Pow(fsigma, 2.0/m.power)
		err := m.errSampler() * math.Sqrt(sigma2)

		residuals[resIdx] = err
		resIdx = (resIdx + 1) % N

		sigmas2[i] = sigma2
		errs[i] = err
	}

	return errs, sigmas2
}
