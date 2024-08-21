package calculations

import (
	"math"
	"sort"
)

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

func FindMedian(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	sort.Float64s(nums)

	if len(nums)%2 == 0 {
		return (nums[len(nums)/2-1] + nums[len(nums)/2]) / 2
	} else {
		return nums[len(nums)/2]
	}
}

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

func FindStandardDeviation(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	return math.Sqrt(FindVariance(nums))
}
