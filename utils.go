package plot

import (
	"fmt"
	"math"
)

func GetMaxFloat64From2dSlice(slices [][]float64) (float64, error) {
	if len(slices) == 0 {
		return 0, fmt.Errorf("cannot get max value from empty slice")
	}
	var max float64
	for _, slice := range slices {
		for _, val := range slice {
			if val > max {
				max = val
			}
		}
	}
	return max, nil
}

func minMaxFloat64Slice(v []float64) (float64, float64) {
	min := math.Inf(1)
	max := math.Inf(-1)

	if len(v) == 0 {
		panic("Empty slice")
	}

	for _, e := range v {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	return min, max
}

func round(input float64) float64 {
	if math.IsNaN(input) {
		return math.NaN()
	}
	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}
	_, decimal := math.Modf(input)
	var rounded float64
	if decimal >= 0.5 {
		rounded = math.Ceil(input)
	} else {
		rounded = math.Floor(input)
	}
	return rounded * sign
}
