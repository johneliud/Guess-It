package readstandardinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/johneliud/Guess-It/guessrange"
)


// ReadStandardInput reads numbers from the standard input and prints the range in which the next number is likely to be.
func ReadStandardInput() {
	var nums []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			continue
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("Failed parsing input to float: %v\n", err)
			continue
		}
		nums = append(nums, num)

		// Calculate the range only if the length of nums is greater than 1
		if len(nums) > 1 {
			// Trim the length of nums to be at most 5 by removing the first element
			if len(nums) == 5 {
				nums = nums[1:]
			}
			lowerLimit, upperLimit := guessrange.GuessRange(nums)

			fmt.Printf("%d %d\n", lowerLimit, upperLimit)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Failed reading input: %v\n", err)
		}
	}
}
