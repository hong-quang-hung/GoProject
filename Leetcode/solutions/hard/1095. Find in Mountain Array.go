package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/find-in-mountain-array/
func init() {
	Solutions[1095] = func() {
		fmt.Println("Input: array = [1,2,3,4,5,3,1], target = 3")
		fmt.Println("Output:", findInMountainArray(3, MountainArrayConstructor([]int{1, 2, 3, 4, 5, 3, 1})))
		fmt.Println("Input: array = [0,1,2,4,2,1], target = 3")
		fmt.Println("Output:", findInMountainArray(3, MountainArrayConstructor([]int{0, 1, 2, 4, 2, 1})))
		fmt.Println("Input: array = [0,5,3,1], target = 1")
		fmt.Println("Output:", findInMountainArray(1, MountainArrayConstructor([]int{0, 5, 3, 1})))
	}
}

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
	arr []int
}

func (m *MountainArray) get(index int) int {
	return m.arr[index]
}
func (m *MountainArray) length() int {
	return len(m.arr)
}

func MountainArrayConstructor(array []int) *MountainArray {
	return &MountainArray{arr: array}
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	left, right := 0, n-1
	mid := (right + left) / 2
	for mountainArr.get(mid-1) >= mountainArr.get(mid) || mountainArr.get(mid) <= mountainArr.get(mid+1) {
		if mountainArr.get(mid) <= mountainArr.get(mid+1) {
			left = mid
		} else {
			right = mid
		}
		mid = (right + left) / 2
	}

	if mountainArr.get(mid) == target {
		return mid
	}

	res := sort.Search(mid, func(i int) bool {
		return mountainArr.get(i) >= target
	})

	if mountainArr.get(res) != target {
		res = sort.Search(n-mid-1, func(i int) bool {
			return mountainArr.get(mid+i+1) <= target
		})
		res = mid + res + 1
		if n == res || mountainArr.get(res) != target {
			return -1
		}
		return res
	}
	return res
}
