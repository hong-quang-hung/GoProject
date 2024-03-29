package easy

import "fmt"

// Reference: https://leetcode.com/problems/add-digits/
func init() {
	Solutions[258] = func() {
		fmt.Println(`Input: num = 38`)
		fmt.Println(`Output:`, addDigits(38))
		fmt.Println(`Input: num = 38`)
		fmt.Println(`Output:`, addDigits(578))
	}
}

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for num > 0 {
			sum += (num % 10)
			num /= 10
		}
		num = sum
	}
	return num
}
