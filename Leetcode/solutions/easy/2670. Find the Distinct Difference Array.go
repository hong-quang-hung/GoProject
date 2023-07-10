package easy

import "fmt"

func init() {
	Solutions[2670] = Leetcode_Distinct_Difference_Array
}

// Reference: https://leetcode.com/problems/find-the-distinct-difference-array/
func Leetcode_Distinct_Difference_Array() {
	fmt.Println("Input: nums = [1,2,3,4,5]")
	fmt.Println("Output:", distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println("Input: nums = [3,2,3,4,2]")
	fmt.Println("Output:", distinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	m1, m2 := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		m2[nums[i]]++
	}

	for i := 0; i < n; i++ {
		m1[nums[i]]++
		m2[nums[i]]--
		if m2[nums[i]] == 0 {
			delete(m2, nums[i])
		}
		res[i] = len(m1) - len(m2)
	}
	return res
}
