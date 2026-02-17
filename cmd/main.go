package main

import (
	"fmt"
	"log"
	"time"

	optimization "github.com/greyskp/cost_optimization/optimizer"
)

type PrintObserver struct{}

func (PrintObserver) Observe(s optimization.Stats) {
	// In real systems: export to metrics/logs/traces, not fmt.Println.
	fmt.Printf(
		"opt n=%d selected=%d leftToFill=%d replacements=%d duration=%s\n",
		s.N, s.SelectedCount, s.LeftToFill, s.Replacements, s.Duration.Truncate(time.Microsecond),
	)
}

func main() {
	costs := []float64{
		-10.0, 20.9, 15.7, 12.4, -40.0, 40.2, 4.7, 60.8, 12.3, -7.6,
		27.6, 4.1, 18.9, 22.7, 31.4, 15.6, 6.2, 29.8, 24.5, 8.9}
	optimized, err := optimization.CostOptimization(costs, optimization.WithObserver(PrintObserver{}))

	if err != nil {
		log.Fatalf("Error detected in CostOptimization: %v", err)
	}

	fmt.Println(optimized)

	total, err := optimization.TotalCost(costs, optimized)
	if err != nil {
		log.Fatalf("Error detected in TotalCost: %v", err)
	}

	fmt.Println(total)
	fmt.Println("===================")

	costsLarge := []float64{
		21.4, 18.6, 33.9, -7.2, 14.8, -26.5, 40.1, -11.3, 9.7, 35.8,
		27.6, -4.1, 18.9, -22.7, 31.4, -15.6, 6.2, 29.8, 24.5, 8.9,
		37.2, -13.4, 11.0, -41.6, 28.1, 17.9, 4.8, -32.5, 19.6, -10.2,
		34.7, -6.5, 16.3, 25.1, 42.0, -14.7, 7.9, -38.4, 22.8, -9.6,
		30.5, 20.3, 12.7, -27.4, 39.1, -5.8, 17.2, -33.9, 26.0, 11.8,
		44.6, 16.5, 8.3, -36.2, 21.9, -7.4, 29.8, -23.6, 14.1, 40.9,
		35.0, -12.1, 10.6, 28.7, 18.4, 6.9, 41.3, 19.2, 25.7, -9.1,
		32.6, 15.4, 5.0, -34.1, 20.3, -8.6, 38.8, -17.0, 13.5, 26.8,
		43.9, 11.6, 16.8, -30.4, 24.1, -7.8, 28.9, -21.7, 9.4, 39.5,
		34.2, -14.9, 6.7, -27.1, 19.8, -10.5, 36.4, -18.3}

	optimizedLarge, err := optimization.CostOptimization(costsLarge, optimization.WithObserver(PrintObserver{}))

	if err != nil {
		log.Fatalf("Error detected in CostOptimization: %v", err)
	}

	fmt.Println(optimizedLarge)

	totalLarge, err := optimization.TotalCost(costsLarge, optimizedLarge)
	if err != nil {
		log.Fatalf("Error detected in TotalCost: %v", err)
	}

	fmt.Println(totalLarge)
	fmt.Println("===================")


	costsNegatives := []float64{
		-21.4, -8.6, -33.9, -7.2, -14.8, -26.5, -40.1, -11.3, -9.7, -35.8,
		-27.6, -4.1, -18.9, -22.7, -31.4, -15.6, -6.2, -29.8, -24.5, -8.9}

	optimizedNegatives, err := optimization.CostOptimization(costsNegatives, optimization.WithObserver(PrintObserver{}))

	if err != nil {
		log.Fatalf("Error detected in CostOptimization: %v", err)
	}

	fmt.Println(optimizedNegatives)

	totalNegatives, err := optimization.TotalCost(costsNegatives, optimizedNegatives)
	if err != nil {
		log.Fatalf("Error detected in TotalCost: %v", err)
	}

	fmt.Println(totalNegatives)
	fmt.Println("===================")


	costsPositives := []float64{
		21.4, 8.6, 33.9, 7.2, 14.8, 26.5, 40.1, 11.3, 9.7, 35.8,
		27.6, 4.1, 18.9, 22.7, 31.4, 15.6, 6.2, 29.8, 24.5, 8.9}

	optimizedPositives, err := optimization.CostOptimization(costsPositives, optimization.WithObserver(PrintObserver{}))

	if err != nil {
		log.Fatalf("Error detected in CostOptimization: %v", err)
	}

	fmt.Println(optimizedPositives)

	totalPositives, err := optimization.TotalCost(costsPositives, optimizedPositives)
	if err != nil {
		log.Fatalf("Error detected in TotalCost: %v", err)
	}

	fmt.Println(totalPositives)

}
