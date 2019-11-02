package euclidean

import (
	"math"
	"metric/internal/testdata"
	"sort"
	"testing"
)

var (
	ai  = testdata.A
	bi  = testdata.B
	min = 0.01
	max = 128
)

func TestMain(m *testing.M) {
	sort.Float64s(ai[:])
	sort.Float64s(bi[:])
	m.Run()
}

func TestDistance(t *testing.T) {
	type args struct {
		a []float64
		b []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{
				a: ai[:],
				b: bi[:],
			},
			want: 0.455,
		},
	}
	for _, tt := range tests {
		// sort.Float64s(tt.args.a)
		// sort.Float64s(tt.args.b)
		t.Run(tt.name, func(t *testing.T) {
			got := Distance(tt.args.a, tt.args.b)
			if !(math.Abs(got-tt.want) < min) {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distance(ai[:], bi[:])
	}

}
