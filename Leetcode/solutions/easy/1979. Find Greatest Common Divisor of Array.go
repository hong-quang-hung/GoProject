package easy

import "fmt"

func init() {
	Solutions[1979] = Leetcode_Find_GCD
}

// Reference: https://leetcode.com/problems/find-greatest-common-divisor-of-array/
func Leetcode_Find_GCD() {
	fmt.Println("Input: nums = [2,5,6,9,10]")
	fmt.Println("Output:", findGCD([]int{2, 5, 6, 9, 10}))
	fmt.Println("Input: nums = [7,5,6,8,3]")
	fmt.Println("Output:", findGCD([]int{7, 5, 6, 8, 3}))
}

func findGCD(nums []int) int {
	minN := nums[0]
	maxN := nums[0]
	for _, v := range nums[1:] {
		if v > maxN {
			maxN = v
		} else if v < minN {
			minN = v
		}
	}
	return gcd(maxN, minN)
}
