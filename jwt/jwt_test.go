package jwt

import (
	"math"
	"testing"
)

var (
	inst = New("ePXXC2v2YCzZFW9yU9Pu2mBc3GgefkEVf5zWhAw9YcvFb8Na")
)

func Benchmark_Decode(b *testing.B) {

	session := Session{
		ID:      math.MaxInt64,
		Expired: math.MaxInt64,
		GroupID: math.MaxUint32,
	}

	str := inst.Encode(&session)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s, ok := inst.Decode(str)
		if !ok {
			b.Errorf("opps: %v %v %v", str, s, ok)
		}
	}
}

func Benchmark_Encode(b *testing.B) {

	session := Session{
		ID:      math.MaxInt64,
		Expired: math.MaxInt64,
		GroupID: math.MaxUint32,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = inst.Encode(&session)
	}
}
