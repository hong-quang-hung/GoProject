package hard

import "fmt"

// Reference: https://leetcode.com/problems/find-in-mountain-array/
func init() {
	Solutions[1095] = func() {
		// fmt.Println("Input: array = [1,2,3,4,5,3,1], target = 3")
		// fmt.Println("Output:", findInMountainArray(3, MountainArrayConstructor([]int{1, 2, 3, 4, 5, 3, 1})))
		// fmt.Println("Input: array = [0,1,2,4,2,1], target = 3")
		// fmt.Println("Output:", findInMountainArray(3, MountainArrayConstructor([]int{0, 1, 2, 4, 2, 1})))
		fmt.Println("Input: array = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,100,99,98,97,96,95,94,93,92,91,90,89,88,87,86,85,84,83,82], target = 101")
		fmt.Println("Output:", findInMountainArray(101, MountainArrayConstructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82})))
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
	left, right := 0, mountainArr.length()-1
	for left <= right {
		mid := (left + right) >> 1
		midVal, leftVal, rightVal := mountainArr.get(mid), mountainArr.get(left), mountainArr.get(right)
		if leftVal == target {
			return left
		}

		fmt.Println(mountainArr.arr[left:right+1], midVal)

		if midVal > target && midVal <= leftVal || midVal < target && midVal > leftVal {
			left = mid + 1
		} else if midVal > target && midVal <= rightVal || midVal < target && midVal >= rightVal {
			left++
			right = mid - 1
		} else {
			left++
		}
	}
	return -1
}
