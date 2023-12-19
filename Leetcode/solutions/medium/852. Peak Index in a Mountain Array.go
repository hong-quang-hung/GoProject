package medium

import "fmt"

// Reference: https://leetcode.com/problems/peak-index-in-a-mountain-array/
func init() {
	Solutions[852] = func() {
		fmt.Println(`Input: arr = [3,4,5,1]`)
		fmt.Println(`Output:`, peakIndexInMountainArray([]int{3, 4, 5, 1}))
		fmt.Println(`Input: arr = [0,2,1,0]`)
		fmt.Println(`Output:`, peakIndexInMountainArray([]int{0, 2, 1, 0}))
		fmt.Println(`Input: arr = [0,10,5,2]`)
		fmt.Println(`Output:`, peakIndexInMountainArray([]int{0, 10, 5, 2}))
	}
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	res := (right + left) / 2
	for arr[res-1] >= arr[res] || arr[res] <= arr[res+1] {
		if arr[res] <= arr[res+1] {
			left = res
		} else {
			right = res
		}
		res = (right + left) / 2
	}
	return res
}
