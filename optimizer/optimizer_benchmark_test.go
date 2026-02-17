package optimization

import (
	"math/rand/v2"
	"testing"
)

var benchOutput []int
var benchError error

func BenchmarkOptimizer_10(b *testing.B) {
	n := 10

	b.Run("Mixed", func(b *testing.B) {
		costs := randFloats(-100.0, 500.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Positives", func(b *testing.B) {
		costs := randFloats(0.0, 500.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Negatives", func(b *testing.B) {
		costs := randFloats(-500.0, -1.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Equals", func(b *testing.B) {
		costs := randFloats(0.0, 0.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

}

func BenchmarkOptimizer_100(b *testing.B) {
	n := 100

	b.Run("Mixed", func(b *testing.B) {
		costs := randFloats(-100.0, 500.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Positives", func(b *testing.B) {
		costs := randFloats(0.0, 500.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Negatives", func(b *testing.B) {
		costs := randFloats(-500.0, -1.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Equals", func(b *testing.B) {
		costs := randFloats(0.0, 0.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

}

func BenchmarkOptimizer_10000(b *testing.B) {
	n := 10000

	b.Run("Mixed", func(b *testing.B) {
		costs := randFloats(-100.0, 500.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Positives", func(b *testing.B) {
		costs := randFloats(0.0, 500.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Negatives", func(b *testing.B) {
		costs := randFloats(-500.0, -1.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

	b.Run("Equals", func(b *testing.B) {
		costs := randFloats(0.0, 0.0, n)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			benchOutput, benchError = CostOptimization(costs)
		}
	})

}

// Helper function to generate lists of inputs
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
