package cos

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

func TestSimilarity(t *testing.T) {
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
			want: 0.99,
		},
		{
			name: "2",
			args: args{
				a: []float64{1, 1, 2, 1.0, 1, 1, 0, 0, 0},
				b: []float64{1, 1, 1, 0, 1, 1, 1, 1, 1},
			},
			want: 0.71,
		},
		{
			name: "3",
			args: args{
				a: []float64{5, 0, 1, 3},
				b: []float64{6, 0, 0, 2},
			},
			want: 0.962,
		},
	}
	for _, tt := range tests {
		// sort.Float64s(tt.args.a)
		// sort.Float64s(tt.args.b)
		t.Run(tt.name, func(t *testing.T) {
			got := Similarity(tt.args.a, tt.args.b)
			if !(math.Abs(got-tt.want) < min) {
				t.Errorf("Similarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Similarity(ai[:], bi[:])
	}

}
