package medium

import "fmt"

// Reference: https://leetcode.com/problems/k-radius-subarray-averages/
func Leetcode_Get_Averages() {
	fmt.Println("Input: nums = [7,4,3,9,1,8,5,2,6], k = 3")
	fmt.Println("Output:", getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	fmt.Println("Input: nums = [100000], k = 0")
	fmt.Println("Output:", getAverages([]int{100000}, 0))
	fmt.Println("Input: nums = [8], k = 100000")
	fmt.Println("Output:", getAverages([]int{8}, 100000))
}

func getAverages(nums []int, k int) []int {
	n, div := len(nums), 2*k+1
	sum := 0
	for i := 0; i < n && i < div; i++ {
		sum += nums[i]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	for i := div - 1; i < n; i++ {
		if i != div-1 {
			sum = sum - nums[i-div] + nums[i]
		}
		res[i-k] = sum / div
	}
	return res
}
