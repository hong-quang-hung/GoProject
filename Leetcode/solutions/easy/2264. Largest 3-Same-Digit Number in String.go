package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/largest-3-same-digit-number-in-string/
func init() {
	Solutions[2264] = func() {
		fmt.Println("Input: num = \"6777133339\"")
		fmt.Println("Output:", largestGoodInteger("6777133339"))
		fmt.Println("Input: num = \"2300019\"")
		fmt.Println("Output:", largestGoodInteger("2300019"))
		fmt.Println("Input: num = \"42352338\"")
		fmt.Println("Output:", largestGoodInteger("42352338"))
	}
}

func largestGoodInteger(num string) string {
	res := ""
	for i, j := 0, 1; j < len(num); {
		if num[j] == num[j-1] {
			if j-i == 2 {
				if res < num[i:j+1] {
					res = num[i : j+1]
				}
				i = j
			}
		} else {
			i = j
		}
		j++
	}
	return res
}
