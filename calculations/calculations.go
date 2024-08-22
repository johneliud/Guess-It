package calculations

import (
	"math"
	"sort"
)

// FindAverage finds the average of a given slice of float64 numbers.
func FindAverage(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

// FindMedian finds the median of a given slice of float64 numbers.
func FindMedian(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	// Sort the slice in ascending order
	sort.Float64s(nums)

	if len(nums)%2 == 0 {
		firstMiddleIndex := len(nums)/2 - 1
		secondMiddleIndex := firstMiddleIndex + 1
		return (nums[firstMiddleIndex] + nums[secondMiddleIndex]) / 2
	} else {
		return nums[len(nums)/2]
	}
}

// FindVariance calculates the variance of a given slice of float64 numbers.
func FindVariance(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	mean := FindAverage(nums)
	variance := 0.0

	for _, num := range nums {
		result := (num - mean) * (num - mean)
		variance += result
	}
	return variance / float64(len(nums))
}

// FindStandardDeviation calculates the standard deviation of a given slice of float64 numbers by taking the square root of the variance.
func FindStandardDeviation(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	return math.Sqrt(FindVariance(nums))
}
