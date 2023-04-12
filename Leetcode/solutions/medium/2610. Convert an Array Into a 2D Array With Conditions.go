package medium

import "fmt"

// Reference: https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/
func Leetcode_Find_Matrix() {
	fmt.Println("Input: nums = [1,3,4,1,2,3,1]")
	fmt.Println("Output:", findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output:", findMatrix([]int{1, 2, 3, 4}))
}

func findMatrix(nums []int) [][]int {
	res := make([][]int, 1)
	m := make(map[int]int)
	for _, num := range nums {
		if count, c := m[num]; c {
			if len(res) < count+1 {
				res = append(res, []int{})
			}
			res[count] = append(res[count], num)
			m[num] = count + 1
		} else {
			res[0] = append(res[0], num)
			m[num] = 1
		}
	}
	return res
}
