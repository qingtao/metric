package cos

import "math"

// Similarity 相似度
func Similarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return math.MaxFloat64
	}
	if len(a) == 0 {
		return 0
	}
	var (
		sum  float64
		sumA float64
		sumB float64
	)
	for i := 0; i < len(a); i++ {
		sumA += a[i] * a[i]
		sumB += b[i] * b[i]
		sum += a[i] * b[i]
	}
	return sum / (math.Sqrt(sumA) * math.Sqrt(sumB))
}
