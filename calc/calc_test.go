package calc_test

import (
	"testing"

	"github.com/LarsFox/btc-ticker/calc"
	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	vals := []float64{4.5, 0, 4.5, 5.5, 5.5}
	assert.EqualValues(t, 4, calc.Mean(vals))
}

func TestMeannEmpty(t *testing.T) {
	assert.EqualValues(t, 0, calc.Mean([]float64{}))
}

func TestMedianEven(t *testing.T) {
	vals := []float64{7, 9, 3, 2, 1, 17, 0, 228, 36.6, 100500}
	assert.EqualValues(t, 8, calc.Median(vals))
}

func TestMedianOdd(t *testing.T) {
	vals := []float64{7, 3, 2, 1, 17, 0, 228, 36.6, 100500}
	assert.EqualValues(t, 7, calc.Median(vals))
}

func TestMedianEmpty(t *testing.T) {
	assert.EqualValues(t, 0, calc.Median([]float64{}))
}
