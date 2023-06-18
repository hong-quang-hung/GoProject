package medium

import "fmt"

// Reference: https://leetcode.com/problems/permutations/
func Leetcode_Permute() {
	fmt.Println("Input: nums = [1,2,3]")
	fmt.Println("Output:", permute([]int{1, 2, 3}))
	fmt.Println("Input: nums = [0,1]")
	fmt.Println("Output:", permute([]int{0, 1}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	permuteRS(&res, nums, []int{}, make([]bool, len(nums)))
	return res
}

func permuteRS(res *[][]int, nums []int, current []int, visit []bool) {
	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
		return
	}

	for j := 0; j < len(nums); j++ {
		if !visit[j] {
			visit[j] = true
			current = append(current, nums[j])
			permuteRS(res, nums, current, visit)
			current = current[:len(current)-1]
			visit[j] = false
		}
	}
}
