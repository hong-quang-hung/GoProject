package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-unique-binary-string/
func init() {
	Solutions[1980] = func() {
		fmt.Println("Input: nums = [\"01\",\"10\"]")
		fmt.Println("Output:", findDifferentBinaryString([]string{"01", "10"}))
		fmt.Println("Input: nums = [\"00\",\"01\"]")
		fmt.Println("Output:", findDifferentBinaryString([]string{"00", "01"}))
		fmt.Println("Input: nums = [\"111\",\"011\",\"001\"]")
		fmt.Println("Output:", findDifferentBinaryString([]string{"111", "011", "001"}))
	}
}

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	res, m := make([]rune, n), []rune{'1', '0'}
	for i, num := range nums {
		res[i] = m[num[i]-'0']
	}
	return string(res)
}
