package euclidean

import (
	"math"
)

// Distance 距离
func Distance(a, b []float64) float64 {
	if len(a) != len(b) {
		return 1.0
	}
	if len(a) == 0 {
		return 1.0
	}
	var (
		sum float64
		tmp float64
	)
	for i := 0; i < len(a); i++ {
		tmp = a[i] - b[i]
		sum += tmp * tmp
	}
	return math.Sqrt(sum)
}
