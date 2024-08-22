package guessrange

import (
	"math"

	"github.com/johneliud/Guess-It/calculations"
)


// GuessRange calculates the range in which the next number is likely to be based on the given set of numbers.
func GuessRange(nums []float64) (int, int) {
	mean := calculations.FindAverage(nums)
	standardDeviation := math.Sqrt(calculations.FindVariance(nums))

	// It uses the Z-Score of 2.0 to have a 95% confidence interval.
	lowerLimit := mean - (2.0 * standardDeviation)
	upperLimit := mean + (2.0 * standardDeviation)
	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}
