package easy

import "fmt"

func init() {
	Solutions[1805] = Leetcode_Num_Different_Integers
}

// Reference: https://leetcode.com/problems/number-of-different-integers-in-a-string/
func Leetcode_Num_Different_Integers() {
	fmt.Println("Input: word = 'a123bc34d8ef34'")
	fmt.Println("Output:", numDifferentIntegers("a123bc34d8ef34"))
	fmt.Println("Input: word = 'leet1234code234'")
	fmt.Println("Output:", numDifferentIntegers("leet1234code234"))
	fmt.Println("Input: word = 'a1b01c001'")
	fmt.Println("Output:", numDifferentIntegers("a1b01c001"))
}

func numDifferentIntegers(word string) int {
	m := make(map[float64]bool)
	num, pow := float64(-1), float64(1)
	for i := len(word) - 1; i >= 0; i-- {
		digit := int(word[i] - '0')
		if digit >= 0 && digit <= 9 {
			if num != -1 {
				num += float64(digit) * pow
			} else {
				num = float64(digit) * pow
			}
			pow *= 10
		} else {
			if num != -1 {
				m[num] = true
			}
			num = -1
			pow = 1
		}
	}

	if num != -1 {
		m[num] = true
	}
	return len(m)
}
