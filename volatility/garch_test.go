package volatility

import (
	"math"
	"math/rand"
	"testing"
)

func TestNewGARCH(t *testing.T) {

	m := NewGARCH(
		2.,
		0.01,
		[]float64{0.1000},
		[]float64{},
		[]float64{0.8750})
	N := 10

	errors := make([]float64, N)
	for i := 0; i < N; i++ {
		errors[i] = rand.NormFloat64()
	}

	i := 0
	errSampler := func() float64 {
		i += 1
		return errors[i-1]
	}
	m.SetErrorSampler(errSampler)

	persistence := 0.5*m.alpha[0] + m.beta[0]
	initialValue := m.omega / (1 - persistence)

	sigmas2 := make([]float64, N)
	data := make([]float64, N)
	for t := 0; t < N; t++ {
		sigmas2[t] = m.omega
		var shock float64
		if t == 0 {
			shock = 0.
		} else {
			shock = data[t-1] * data[t-1]
		}
		sigmas2[t] += m.alpha[0] * shock
		var laggedValue float64
		if t == 0 {
			laggedValue = initialValue
		} else {
			laggedValue = sigmas2[t-1]
		}
		sigmas2[t] += m.beta[0] * laggedValue
		data[t] = errors[t] * math.Sqrt(sigmas2[t])
	}

	sampleData, sampleSigmas2 := m.Sample(N)
	for i := 0; i < N; i++ {
		if math.Abs(sampleData[i]-data[i]) > 0.00001 {
			t.Fatalf("different data from sampling: %f %f", sampleData[i], data[i])
		}
		if math.Abs(sampleSigmas2[i]-sigmas2[i]) > 0.00001 {
			t.Fatalf("different sigma from sampling: %f %f", sampleSigmas2[i], sigmas2[i])
		}
	}
}

func TestGARCH_Progress(t *testing.T) {

	m := NewGARCH(
		2.,
		0.01,
		[]float64{0.1000},
		[]float64{},
		[]float64{0.8750})
	N := 10

	errors := make([]float64, N)
	for i := 0; i < N; i++ {
		errors[i] = rand.NormFloat64()
	}

	i := 0
	errSampler := func() float64 {
		i += 1
		return errors[i-1]
	}
	m.SetErrorSampler(errSampler)

	persistence := 0.5*m.alpha[0] + m.beta[0]
	initialValue := m.omega / (1 - persistence)

	sigmas2 := make([]float64, 2*N)
	data := make([]float64, 2*N)
	for t := 0; t < N; t++ {
		sigmas2[t] = m.omega
		var shock float64
		if t == 0 {
			shock = 0.
		} else {
			shock = data[t-1] * data[t-1]
		}
		sigmas2[t] += m.alpha[0] * shock
		var laggedValue float64
		if t == 0 {
			laggedValue = initialValue
		} else {
			laggedValue = sigmas2[t-1]
		}
		sigmas2[t] += m.beta[0] * laggedValue
		data[t] = rand.NormFloat64() * math.Sqrt(sigmas2[t])
	}

	for i := 0; i < N; i++ {
		m.Progress(data[i])
	}

	for t := N; t < 2*N; t++ {
		sigmas2[t] = m.omega
		shock := data[t-1] * data[t-1]
		sigmas2[t] += m.alpha[0] * shock
		var laggedValue float64
		if t == N {
			laggedValue = initialValue
		} else {
			laggedValue = sigmas2[t-1]
		}
		sigmas2[t] += m.beta[0] * laggedValue
		data[t] = errors[t-N] * math.Sqrt(sigmas2[t])
	}

	sampleData, sampleSigmas2 := m.Sample(N)
	for i := 0; i < N; i++ {
		if math.Abs(sampleSigmas2[i]-sigmas2[N+i]) > 0.00001 {
			t.Fatalf("different sigma from sampling: %f %f", sampleSigmas2[i], sigmas2[i])
		}
		if math.Abs(sampleData[i]-data[N+i]) > 0.00001 {
			t.Fatalf("different data from sampling: %f %f", sampleData[i], data[i])
		}
	}
}
