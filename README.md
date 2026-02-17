# Cost Optimization Library
## Overview

This project implements a small cost optimization library designed to be used as part of a larger system.

Given a list of costs, the library produces a binary output vector (0 or 1) that:

- minimizes the total cost

- ensures at least half of the outputs are 1.

Total cost is defined as:
```
sum(output[i] * cost[i])
```

Example:
```
costs:   [20.0, -10.0]
output:  [0, 1]
total:   -10.0
```

## API
```
CostOptimization(costs []float64, opts ...Option) ([]int, error)
TotalCost(costs []float64, optimization []int) (float64, error)
```

### Inputs

- costs: list of real numbers

### Outputs

- binary slice ([]int) of same length
- values are 0 or 1

## Assumptions

The following assumptions were made to clarify unspecified behavior:

- Costs are represented as float64.

- NaN values are invalid and return an error.

- +Inf and -Inf are allowed.

- The constraint is interpreted as at least ⌈n/2⌉ selected elements.

- If more than ⌈n/2⌉ costs are negative, all negative costs are selected.

- Multiple optimal solutions may exist; a deterministic tie-break rule is used:

- For equal costs, lower indices are preferred.

- The library does not perform logging or exit the process.

- The input slice fits in memory.

- Typical input size is around ~100 elements.

## Decision Logic
### Key observation

Selecting a negative cost always decreases the total cost, so:

All negative costs should be selected.

If fewer than ⌈n/2⌉ elements are selected:

- choose the smallest remaining costs until the constraint is met.

### Algorithm

Scan input once:

- validate values

- select all negative costs.

Compute how many additional items must be selected.

Use a fixed-size max-heap to track the best remaining candidates:

- heap size = number of missing selections.

- the heap top represents the worst currently selected candidate.

Replace heap top whenever a better candidate is found.

Mark selected indices in the output.

### Complexity

- Time: O(n log k)
where k = ceil(n/2) - negatives

- Worst case: O(n log n)

- Space: O(n) (output + heap)

## Edge Cases Handled

- Empty input → error

- NaN values → error

- All negative values → fast path (all selected)

- All positive values

- All equal values (deterministic output)

- Odd and even input lengths

- Very large positive / negative numbers

## Monitoring and Performance
### Monitoring Design

This is a library, so it does not log or export metrics directly.

Instead, an optional Observer hook can be injected:

```
CostOptimization(costs, WithObserver(myObserver))
```

The observer receives a per-call Stats structure containing:

- N (input size)

- SelectedCount

- LeftToFill

- Replacements (heap replacements)

- Duration

This allows the calling system to:

- export metrics to Prometheus/OpenTelemetry

- compute p95/p99 latency externally

- trigger alerts without coupling monitoring to the library.

A NoOpObserver is used by default.

## Performance Benchmarks

Benchmarks were run using:
```
go test -bench . -benchmem ./...
```

Environment:

- Apple M2

- macOS (arm64)

- Go benchmark framework

Example results:

|   Case              |     Time        |     Allocations   |
|     :---:           |    :----:       |       :---:       |
| n=100 (Mixed)       | ~1.75 µs/op     | 38 allocs         |
| n=100 (Positives)   | ~2.73 µs/op     | 60 allocs         |
| n=100 (Negatives)   | ~0.20 µs/op     | 2 allocs          |
| n=10000 (Positives) | ~611.69 µs/op   | 5010 allocs       |


### Interpretation

- Negatives-heavy inputs trigger a fast path (no heap required).

- Heap usage dominates runtime when many positive values exist.

- Equal-value inputs exercise tie-breaking logic but remain deterministic.

## Example Usage
```
type PrintObserver struct{}

func (PrintObserver) Observe(s optimization.Stats) {
    fmt.Printf("n=%d duration=%s\n", s.N, s.Duration)
}

result, err := optimization.CostOptimization(
    costs,
    optimization.WithObserver(PrintObserver{}),
)
```

## One Improvement With More Time

If more time were available, the next improvement would be:

Reduce allocations and improve scalability

- use a linear-time selection algorithm (k-th smallest) instead of heap maintenance

- preallocate heap memory to reduce allocations

- add property-based tests for optimality guarantees

- add optional context support for cancellation in large workloads

## Design Considerations

This solution focuses on:

- clear API boundaries

- deterministic behavior

- strong input validation

- library-safe error handling

- extensibility through options and observer hooks

## Running Tests and Benchmarks
```
go test ./optimizer
go test -bench . -benchmem ./optimizer
```
