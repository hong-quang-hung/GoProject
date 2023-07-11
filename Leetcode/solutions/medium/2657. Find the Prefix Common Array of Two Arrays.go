package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/
func init() {
	Solutions[2657] = func() {
		fmt.Println("Input: A = [1,3,2,4], B = [3,1,2,4]")
		fmt.Println("Output:", findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
		fmt.Println("Input: A = [2,3,1], B = [3,1,2]")
		fmt.Println("Output:", findThePrefixCommonArray([]int{2, 3, 1}, []int{3, 1, 2}))
	}
}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	res := make([]int, n)
	mA, mB := make(map[int]bool), make(map[int]bool)

	mA[A[0]] = true
	mB[B[0]] = true
	if A[0] == B[0] {
		res[0] = 1
	}

	for i := 1; i < n; i++ {
		if A[i] != B[i] {
			if mA[B[i]] {
				res[i]++
			}

			if mB[A[i]] {
				res[i]++
			}
		} else {
			res[i]++
		}

		res[i] += res[i-1]
		mA[A[i]] = true
		mB[B[i]] = true
	}
	return res
}
