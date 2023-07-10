package easy

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[703] = Leetcode_Kth_Largest
}

// Reference: https://leetcode.com/problems/kth-largest-element-in-a-stream/
func Leetcode_Kth_Largest() {
	fmt.Println("Input:")
	fmt.Println("['KthLargest', 'add', 'add', 'add', 'add', 'add']")
	fmt.Println("[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]")
	fmt.Println("Output:")

	kthLargest := KthLargestConstructor(3, []int{4, 5, 8, 2})
	fmt.Println("kthLargest.Add(3)  -->", kthLargest.Add(3))  // return 4
	fmt.Println("kthLargest.Add(5)  -->", kthLargest.Add(5))  // return 5
	fmt.Println("kthLargest.Add(10) -->", kthLargest.Add(10)) // return 5
	fmt.Println("kthLargest.Add(9)  -->", kthLargest.Add(9))  // return 8
	fmt.Println("kthLargest.Add(4)  -->", kthLargest.Add(4))  // return 8
}

type KthLargest struct {
	k    int
	nums []int
}

func KthLargestConstructor(k int, nums []int) KthLargest {
	kth := new(KthLargest)
	kth.k = k
	kth.nums = nums
	sort.Ints(kth.nums)
	return *kth
}

func (kth *KthLargest) Add(val int) int {
	index := sort.SearchInts(kth.nums, val)
	if index == 0 {
		kth.nums = append([]int{val}, kth.nums...)
	} else if index == len(kth.nums) {
		kth.nums = append(kth.nums, val)
	} else {
		kth.nums = append(kth.nums[:index+1], kth.nums[index:]...)
	}
	kth.nums[index] = val
	return kth.nums[len(kth.nums)-kth.k]
}
