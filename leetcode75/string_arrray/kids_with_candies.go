package string_arrray

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := slices.Max(candies)
	candiesValid := maxCandies - extraCandies
	n := len(candies)
	x := make([]bool, n)

	for i := 0; i < n; i++ {
		x[i] = candies[i] >= candiesValid
	}

	return x
}
