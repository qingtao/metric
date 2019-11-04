package euclidean

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

func TestDistance(t *testing.T) {
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
			want: 4.83,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distance(tt.args.a, tt.args.b)
			if !(float32(math.Abs(float64(got-tt.want))) < min) {
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
