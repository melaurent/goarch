package mean

type ConstantMeanProcess struct {
	mean float64
}

func NewConstantMeanProcess(mean float64) *ConstantMeanProcess {
	return &ConstantMeanProcess{mean: mean}
}

func (m *ConstantMeanProcess) Progress(x float64) float64 {
	return x - m.mean
}

func (m *ConstantMeanProcess) Sample(nIntervals int) []float64 {
	samples := make([]float64, nIntervals, nIntervals)
	for i := 0; i < nIntervals; i++ {
		samples[i] = m.mean
	}

	return samples
}
