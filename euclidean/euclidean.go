package euclidean

import (
	"math"
)

// Distance 距离
func Distance(a, b []float32) float32 {
	if len(a) != len(b) {
		return math.MaxFloat32
	}
	if len(a) == 0 {
		return 0
	}
	var (
		sum float32
		tmp float32
	)
	for i := 0; i < len(a); i++ {
		tmp = a[i] - b[i]
		sum += tmp * tmp
	}
	return float32(math.Sqrt(float64(sum)))
}
