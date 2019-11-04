package cos

import "math"

// Similarity 相似度
func Similarity(a, b []float32) float32 {
	if len(a) != len(b) {
		return math.MaxFloat32
	}
	if len(a) == 0 {
		return 0
	}
	var (
		sum  float32
		sumA float32
		sumB float32
	)
	for i := 0; i < len(a); i++ {
		sumA += a[i] * a[i]
		sumB += b[i] * b[i]
		sum += a[i] * b[i]
	}
	return sum / float32(math.Sqrt(float64(sumA)*math.Sqrt(float64(sumB))))
}
