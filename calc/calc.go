package calc

import "sort"

func Mean(f []float64) float64 {
	if len(f) == 0 {
		return 0
	}

	var sum float64
	for _, price := range f {
		sum += price
	}

	return sum / float64(len(f))
}

func Median(f []float64) float64 {
	if len(f) == 0 {
		return 0
	}

	sorted := make([]float64, len(f))
	copy(sorted, f)

	sort.Float64s(sorted)

	if len(f)%2 == 0 {
		return (sorted[len(f)/2-1] + sorted[len(f)/2]) / 2
	}

	return sorted[len(f)/2]
}
