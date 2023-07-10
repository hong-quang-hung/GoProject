package easy

import "fmt"

func init() {
	Solutions[989] = Leetcode_Add_To_Array_Form
}

// Reference: https://leetcode.com/problems/add-to-array-form-of-integer/
func Leetcode_Add_To_Array_Form() {
	fmt.Println("Input: num = [1,2,0,0], k = 34")
	fmt.Println("Output:", addToArrayForm([]int{1, 2, 0, 0}, 34))
}

func addToArrayForm(num []int, k int) []int {
	var ext int
	n := len(num)
	i := n - 1
	for k > 0 {
		if i >= 0 {
			ext = (num[i] + k%10) / 10
			num[i] = (num[i] + k%10) - ext*10
		} else {
			ext = 0
			num = append([]int{k % 10}, num...)
		}
		k = k/10 + ext
		i--
	}
	return num
}
