package optimization

import (
	"container/heap"
	"errors"
	"math"
	"time"
)

var ErrEmptyInput = errors.New("input costs slice is empty")
var ErrInvalidNumber = errors.New("NaN error, not a valid number")
var ErrDifferentSizes = errors.New("The length of the costs and the output are different")

type cost struct {
	price float64
	index int
}

type MaxHeap []cost

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return (h[i].price > h[j].price) || (h[i].price == h[j].price && h[i].index > h[j].index)
}                               // Use of '>' to provide the maximum instead of the default minimum
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(cost)) }

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// CostOptimization returns a binary slice indicating which prices should be selected to minimize total cost ensuring at least half of the input prices are selected, prioritizing negative values and the smallest positive values.
func CostOptimization(prices []float64, opts ...Option) ([]int, error) {

	cfg := applyOptions(opts)

	start := time.Now()
	var selectedCount int
	var leftToFill int
	var replacements int

	// Ensure we always emit stats once, even on early returns/errors.
	defer func() {
		cfg.observer.Observe(Stats{
			N:             len(prices),
			SelectedCount: selectedCount,
			LeftToFill:    leftToFill,
			Replacements:  replacements,
			Duration:      time.Since(start),
		})
	}()

	if len(prices) == 0 {
		return nil, ErrEmptyInput
	}

	// Number of elements to be added to reach at least n/2
	minSize := 0
	if len(prices)%2 == 0 {
		minSize = len(prices) / 2
	} else {
		minSize = len(prices)/2 + 1
	}

	res := make([]int, len(prices))
	//count := 0

	for i, value := range prices {
		if math.IsNaN(value) {
			return nil, ErrInvalidNumber
		}
		if value < 0 {
			res[i] = 1
			selectedCount++
		}
	}

	// If there is enough negative costs return only those
	if selectedCount >= minSize {
		leftToFill = 0
		return res, nil
	}

	// Initialize the heap used to keep track of the highest element
	smallest := &MaxHeap{}
	heap.Init(smallest)
	leftToFill = minSize - selectedCount

	for index, value := range prices {
		// fill with the first values available
		if res[index] == 0 {
			c := cost{value, index}

			if smallest.Len() < leftToFill {
				heap.Push(smallest, c)
				continue
			}

			highest := (*smallest)[0]
			if value < highest.price || (value == highest.price && index < highest.index) {
				(*smallest)[0] = c
				heap.Fix(smallest, 0)
				replacements++
			}
		}

	}

	// update res with the missing element
	for _, v := range *smallest {
		res[v.index] = 1
		selectedCount++
	}

	return res, nil
}

// TotalCost calculates the total cost by multiplying each price with its corresponding optimization flag and summing the results.
func TotalCost(prices []float64, optimization []int) (float64, error) {
	result := 0.0
	if len(prices) != len(optimization) {
		return 0.0, ErrDifferentSizes
	}
	for i := range prices {
		if prices[i] == math.Inf(1) || prices[i] == math.Inf(-1) {
			if optimization[i] == 1 {
				return prices[i], nil
			}
		} else {
			result += prices[i] * float64(optimization[i])
		}

	}

	return result, nil
}
