package solutions

import "fmt"

func Leetcode_Two_Sum() {
	target := 9
	nums := [...]int{2, 7, 11, 15}
	fmt.Println("Input: nums = [2,7,11,15], target = 9")
	fmt.Println("Output:", twoSum(nums[:], target))
}
