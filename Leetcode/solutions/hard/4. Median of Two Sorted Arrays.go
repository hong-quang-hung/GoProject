package hard

import "fmt"

// Reference: https://leetcode.com/problems/median-of-two-sorted-arrays/
func init() {
	Solutions[4] = func() {
		fmt.Println("Input: nums1 = [1], nums2 = [2,3,4,5,6]")
		fmt.Println("Output:", findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6}))
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	var l int = (m + n + 1) >> 1
	var r int = (m + n + 2) >> 1
	return float64(getkth(nums1, m, nums2, n, l)+getkth(nums1, m, nums2, n, r)) / 2
}

func getkth(s []int, m int, l []int, n int, k int) int {
	if m > n {
		return getkth(l, n, s, m, k)
	}

	if m == 0 {
		return l[k-1]
	}

	if k == 1 {
		return min(s[0], l[0])
	}

	i, j := min(m, k/2), k/2
	if s[i-1] > l[j-1] {
		l = l[j:]
		return getkth(s, m, l, n-j, k-j)
	} else {
		s = s[i:]
		return getkth(s, m-i, l, n, k-i)
	}
}
