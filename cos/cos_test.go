package cos

import (
	"math"
	"metric/internal/testdata"
	"sort"
	"testing"
)

var (
	ai          = testdata.A
	bi          = testdata.B
	min float32 = 0.01
	max float32 = 128
)

func TestMain(m *testing.M) {
	sort.Slice(ai[:], func(i, j int) bool {
		return ai[i] < ai[j]
	})
	sort.Slice(bi[:], func(i, j int) bool {
		return bi[j] < bi[j]
	})
	m.Run()
}

func TestSimilarity(t *testing.T) {
	type args struct {
		a []float32
		b []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "1",
			args: args{
				a: ai[:],
				b: bi[:],
			},
			want: 1.94,
		},
		{
			name: "2",
			args: args{
				a: []float32{1, 1, 2, 1.0, 1, 1, 0, 0, 0},
				b: []float32{1, 1, 1, 0, 1, 1, 1, 1, 1},
			},
			want: 1.189,
		},
		{
			name: "3",
			args: args{
				a: []float32{5, 0, 1, 3},
				b: []float32{6, 0, 0, 2},
			},
			want: 2.41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Similarity(tt.args.a, tt.args.b)
			if !(float32(math.Abs(float64(got-tt.want))) < min) {
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
