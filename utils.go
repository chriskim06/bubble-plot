package plot

import (
	"math"

	"github.com/charmbracelet/lipgloss"
	"github.com/chriskim06/drawille-go"
)

func colorToInt(c lipgloss.Color) drawille.Color {
	r, g, b, _ := c.RGBA()
	if r == g && g == b {
		if r < 8 {
			return drawille.Color(16)
		} else if r > 248 {
			return drawille.Color(231)
		}
		return drawille.Color(int(math.Round((float64(r-8)/247)*24) + 232))
	}
	return drawille.Color(int(16 +
		(36 + math.Round(float64(r/255*5))) +
		(6 * math.Round(float64((g / 255 * 5)))) +
		math.Round(float64(b/255*5))))
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
