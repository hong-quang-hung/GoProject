package medium

import "fmt"

// Reference: https://leetcode.com/problems/neighboring-bitwise-xor/
func Leetcode_Does_Valid_Array_Exist() {
	fmt.Println("Input: derived = [1,1,0]")
	fmt.Println("Output:", doesValidArrayExist([]int{1, 1, 0}))
	fmt.Println("Input: derived = [1,1]")
	fmt.Println("Output:", doesValidArrayExist([]int{1, 1}))
	fmt.Println("Input: derived = [1,0]")
	fmt.Println("Output:", doesValidArrayExist([]int{1, 0}))
}

func doesValidArrayExist(derived []int) bool {
	derived = append(derived, derived[0])
	for choose := 0; choose <= 1; choose++ {
		next := choose
		for i := 1; i < len(derived); i++ {
			next = derived[i] ^ next
		}

		if next == choose {
			return true
		}
	}
	return false
}
