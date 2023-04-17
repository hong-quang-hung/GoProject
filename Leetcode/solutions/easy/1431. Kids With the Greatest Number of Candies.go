package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
func Leetcode_Kids_With_Candies() {
	fmt.Println("Input: candies = [2,3,5,1,3], extraCandies = 3")
	fmt.Println("Output:", kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println("Input: candies = [4,2,1,1,2], extraCandies = 1")
	fmt.Println("Output:", kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println("Input: candies = [12,1,12], extraCandies = 10")
	fmt.Println("Output:", kidsWithCandies([]int{12, 1, 12}, 10))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res, maxCandies := make([]bool, len(candies)), 0
	for _, candy := range candies {
		maxCandies = max(maxCandies, candy)
	}

	for i, candy := range candies {
		res[i] = maxCandies <= extraCandies+candy
	}
	return res
}
