package optimization

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestInts(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Error the following inputs are wrong: \n")
	correct := true

	costs := []float64{-10, 17, 15, -40, 20}
	expected := []int{1, 0, 1, 1, 0}
	result, err := CostOptimization(costs)

	if err != nil {
		t.Fatalf("Error thrown from CostOptimization: %v", err)
	}

	outputValidity := onlyOnesAndZeros(result)

	if !outputValidity {
		t.Fatalf("Error, the output contains other elements than 0 and 1. Output: %v", result)
	}

	selected := countOnes(result)
	need := len(costs) / 2

	if len(costs)%2 != 0 {
		need += 1
	}

	if selected < need {
		t.Fatalf("Error, not enough costs were selected. Got %v and expected %v", selected, need)
	}

	for i := range result {
		if expected[i] != result[i] {
			correct = false
			builder.WriteString(strconv.FormatFloat(costs[i], 'f', -1, 64) + ", expected " + strconv.Itoa(expected[i]) + ", got: " + strconv.Itoa(result[i]) + "\n")
		}
	}

	if !correct {
		t.Fatal(builder.String())
	}
}

func TestFloats(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Error the following inputs are wrong: \n")
	correct := true

	costs := []float64{-10.4, 15.5, 15.7, 0.0, 20.4}
	expected := []int{1, 1, 0, 1, 0}
	result, err := CostOptimization(costs)

	if err != nil {
		t.Fatalf("Error thrown from CostOptimization: %v", err)
	}

	outputValidity := onlyOnesAndZeros(result)

	if !outputValidity {
		t.Fatalf("Error, the output contains other elements than 0 and 1. Output: %v", result)
	}

	selected := countOnes(result)
	need := len(costs) / 2

	if len(costs)%2 != 0 {
		need += 1
	}

	if selected < need {
		t.Fatalf("Error, not enough costs were selected. Got %v and expected %v", selected, need)
	}

	for i := range result {
		if expected[i] != result[i] {
			correct = false
			builder.WriteString(strconv.FormatFloat(costs[i], 'f', -1, 64) + ", expected " + strconv.Itoa(expected[i]) + ", got: " + strconv.Itoa(result[i]) + "\n")
		}
	}

	if !correct {
		t.Fatal(builder.String())
	}
}

func TestInfinity(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Error the following inputs are wrong: \n")
	correct := true

	costs := []float64{math.Inf(-1), 3, math.Inf(1)}
	expected := []int{1, 1, 0}
	result, err := CostOptimization(costs)

	if err != nil {
		t.Fatalf("Error thrown from CostOptimization: %v", err)
	}

	outputValidity := onlyOnesAndZeros(result)

	if !outputValidity {
		t.Fatalf("Error, the output contains other elements than 0 and 1. Output: %v", result)
	}

	selected := countOnes(result)
	need := len(costs) / 2

	if len(costs)%2 != 0 {
		need += 1
	}

	if selected < need {
		t.Fatalf("Error, not enough costs were selected. Got %v and expected %v", selected, need)
	}

	for i := range result {
		if expected[i] != result[i] {
			correct = false
			builder.WriteString(strconv.FormatFloat(costs[i], 'f', -1, 64) + ", expected " + strconv.Itoa(expected[i]) + ", got: " + strconv.Itoa(result[i]) + "\n")
		}
	}

	if !correct {
		t.Fatal(builder.String())
	}
}

func TestLarge(t *testing.T) {

	var builder strings.Builder
	builder.WriteString("Error the following inputs are wrong: ")
	correct := true

	costs := []float64{
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

	expected := []int{
		0, 0, 0, 1, 0, 1, 0, 1, 1, 0,
		0, 1, 0, 1, 0, 1, 1, 0, 0, 1,
		0, 1, 1, 1, 0, 0, 1, 1, 0, 1,
		0, 1, 0, 0, 0, 1, 1, 1, 0, 1,
		0, 0, 1, 1, 0, 1, 0, 1, 0, 1,
		0, 0, 1, 1, 0, 1, 0, 1, 1, 0,
		0, 1, 1, 0, 0, 1, 0, 0, 0, 1,
		0, 0, 1, 1, 0, 1, 0, 1, 1, 0,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 0,
		0, 1, 1, 1, 0, 1, 0, 1}
	result, err := CostOptimization(costs)

	if err != nil {
		t.Fatalf("Error thrown from CostOptimization: %v", err)
	}

	outputValidity := onlyOnesAndZeros(result)

	if !outputValidity {
		t.Fatalf("Error, the output contains other elements than 0 and 1. Output: %v", result)
	}

	selected := countOnes(result)
	need := len(costs) / 2

	if len(costs)%2 != 0 {
		need += 1
	}

	if selected < need {
		t.Fatalf("Error, not enough costs were selected. Got %v and expected %v", selected, need)
	}

	for i := range result {
		if expected[i] != result[i] {
			correct = false
			builder.WriteString(strconv.FormatFloat(costs[i], 'f', -1, 64) + ", expected " + strconv.Itoa(expected[i]) + ", got: " + strconv.Itoa(result[i]) + "\n")
		}
	}

	if !correct {
		t.Fatal(builder.String())
	}
}

func TestOnlyNegatives(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Error the following inputs are wrong: ")
	correct := true

	costs := []float64{-987.4, -684.5, -6450.7, -4156.3, -8.4}
	expected := []int{1, 1, 1, 1, 1}
	result, err := CostOptimization(costs)

	if err != nil {
		t.Fatalf("Error thrown from CostOptimization: %v", err)
	}

	outputValidity := onlyOnesAndZeros(result)

	if !outputValidity {
		t.Fatalf("Error, the output contains other elements than 0 and 1. Output: %v", result)
	}

	selected := countOnes(result)
	need := len(costs) / 2

	if len(costs)%2 != 0 {
		need += 1
	}

	if selected < need {
		t.Fatalf("Error, not enough costs were selected. Got %v and expected %v", selected, need)
	}

	for i := range result {
		if expected[i] != result[i] {
			correct = false
			builder.WriteString(strconv.FormatFloat(costs[i], 'f', -1, 64) + ", expected " + strconv.Itoa(expected[i]) + ", got: " + strconv.Itoa(result[i]) + "\n")
		}
	}

	if !correct {
		t.Fatal(builder.String())
	}
}

func TestOnlyPositives(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Error the following inputs are wrong: ")
	correct := true

	costs := []float64{987.4, 684.5, 6450.7, 4156.3, 8.4}
	expected := []int{1, 1, 0, 0, 1}
	result, err := CostOptimization(costs)

	if err != nil {
		t.Fatalf("Error thrown from CostOptimization: %v", err)
	}

	outputValidity := onlyOnesAndZeros(result)

	if !outputValidity {
		t.Fatalf("Error, the output contains other elements than 0 and 1. Output: %v", result)
	}

	selected := countOnes(result)
	need := len(costs) / 2

	if len(costs)%2 != 0 {
		need += 1
	}

	if selected < need {
		t.Fatalf("Error, not enough costs were selected. Got %v and expected %v", selected, need)
	}

	for i := range result {
		if expected[i] != result[i] {
			correct = false
			builder.WriteString(strconv.FormatFloat(costs[i], 'f', -1, 64) + ", expected " + strconv.Itoa(expected[i]) + ", got: " + strconv.Itoa(result[i]) + "\n")
		}
	}

	if !correct {
		t.Fatal(builder.String())
	}
}

func TestOnlyZeros(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("Error the following inputs are wrong: ")
	correct := true

	costs := []float64{0.0, 0, 0.0, 0, 0}
	expected := []int{1, 1, 1, 0, 0}
	result, err := CostOptimization(costs)

	if err != nil {
		t.Fatalf("Error thrown from CostOptimization: %v", err)
	}

	outputValidity := onlyOnesAndZeros(result)

	if !outputValidity {
		t.Fatalf("Error, the output contains other elements than 0 and 1. Output: %v", result)
	}

	selected := countOnes(result)
	need := len(costs) / 2

	if len(costs)%2 != 0 {
		need += 1
	}

	if selected < need {
		t.Fatalf("Error, not enough costs were selected. Got %v and expected %v", selected, need)
	}

	for i := range result {
		if expected[i] != result[i] {
			correct = false
			builder.WriteString(strconv.FormatFloat(costs[i], 'f', -1, 64) + ", expected " + strconv.Itoa(expected[i]) + ", got: " + strconv.Itoa(result[i]) + "\n")
		}
	}

	if !correct {
		t.Fatal(builder.String())
	}
}

// Testing the different errors returned
func TestEmpty(t *testing.T) {

	costs := []float64{}
	expected := "input costs slice is empty"
	result, err := CostOptimization(costs)

	if result != nil {
		t.Fatalf("Error %v was not handled properly and result is not empty", expected)
	}

	if err.Error() != expected {
		t.Fatalf("Empty test result is %v instead of the error %v", err, expected)
	}

}

func TestInvalidNumbers(t *testing.T) {
	costs := []float64{-157.4, 419.5, -56.7, math.Sqrt(-1), -8.4}
	expected := "NaN error, not a valid number"
	result, err := CostOptimization(costs)

	if result != nil {
		t.Fatalf("Error %v was not handled properly and result is not empty", expected)
	}

	if err.Error() != expected {
		t.Fatalf("Invalid Numvbers test result is %v instead of the error %v", err, expected)
	}
}

// TotalCost tests
func TestTotalCostBasic(t *testing.T) {
	prices := []float64{10.5, 20.0, 15.5}
	optimization := []int{1, 0, 1}
	expected := 26.0

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

func TestTotalCostAllNegatives(t *testing.T) {
	prices := []float64{-5.0, -10.0, -15.0}
	optimization, err := CostOptimization(prices)
	if err != nil {
		t.Fatalf("CostOptimization returned unexpected error: %v", err)
	}

	expected := -30.0

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

func TestTotalCostAllPositives(t *testing.T) {
	prices := []float64{15.0, 17.4, 8.3}
	optimization, err := CostOptimization(prices)
	if err != nil {
		t.Fatalf("CostOptimization returned unexpected error: %v", err)
	}

	expected := 23.3

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

func TestTotalCostMixedWithZeros(t *testing.T) {
	prices := []float64{0.0, 15.5, 0.0}
	optimization, err := CostOptimization(prices)
	if err != nil {
		t.Fatalf("CostOptimization returned unexpected error: %v", err)
	}

	expected := 0.0

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

func TestTotalCostDifferentSizes(t *testing.T) {
	prices := []float64{10.0, 14.0, 0.0}
	optimization := []int{1, 0}
	expected := "The length of the costs and the output are different"

	result, err := TotalCost(prices, optimization)
	if err == nil {
		t.Fatalf("TotalCost should return error for different sizes")
	}
	if result != 0.0 {
		t.Fatalf("TotalCost should return 0.0 on error, got %v", result)
	}
	if err.Error() != expected {
		t.Fatalf("Error message got %v, expected %v", err, expected)
	}
}

func TestTotalCostEmptySlices(t *testing.T) {
	prices := []float64{}
	optimization := []int{}
	expected := 0.0

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

func TestTotalCostLargeValues(t *testing.T) {
	prices := []float64{1e10, 2e10, 3e10}
	optimization, err := CostOptimization(prices)
	if err != nil {
		t.Fatalf("CostOptimization returned unexpected error: %v", err)
	}
	expected := 3e10

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

func TestTotalCostWithPositiveInfinity(t *testing.T) {
	prices := []float64{math.Inf(1), 10.0, -5.0}
	optimization, err := CostOptimization(prices)
	if err != nil {
		t.Fatalf("CostOptimization returned unexpected error: %v", err)
	}
	expected := 5.0

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

func TestTotalCostWithNegativeInfinity(t *testing.T) {
	prices := []float64{math.Inf(-1), 10.0, -5.0}
	optimization, err := CostOptimization(prices)
	if err != nil {
		t.Fatalf("CostOptimization returned unexpected error: %v", err)
	}
	expected := math.Inf(-1)

	result, err := TotalCost(prices, optimization)
	if err != nil {
		t.Fatalf("TotalCost returned unexpected error: %v", err)
	}
	if result != expected {
		t.Fatalf("TotalCost got %v, expected %v", result, expected)
	}
}

// Testing helpers

func countOnes(optimized []int) (res int) {
	for _, value := range optimized {
		if value == 1 {
			res++
		}
	}

	return
}

func onlyOnesAndZeros(optimized []int) bool {
	for _, value := range optimized {
		if value != 1 && value != 0 {
			return false
		}
	}

	return true
}
