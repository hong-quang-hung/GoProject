package easy

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2570] = Leetcode_Merge_Arrays
}

// Reference: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
func Leetcode_Merge_Arrays() {
	fmt.Println("Input: nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]")
	fmt.Println("Output:", mergeArrays(S2SoSliceInt("[[1,2],[2,3],[4,5]]"), S2SoSliceInt("[[1,4],[3,2],[4,1]]")))
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int][]int)
	for _, num := range nums1 {
		m[num[0]] = num
	}

	for _, num := range nums2 {
		if m[num[0]] == nil {
			m[num[0]] = []int{num[0], num[1]}
		} else {
			m[num[0]][1] += num[1]
		}
	}

	res := make([][]int, 0)
	for _, v := range m {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	return res
}
