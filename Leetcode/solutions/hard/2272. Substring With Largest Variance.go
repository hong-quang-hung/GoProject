package hard

import "fmt"

func init() {
	Solutions[2272] = Leetcode_Largest_Variance
}

// Reference: https://leetcode.com/problems/substring-with-largest-variance/
func Leetcode_Largest_Variance() {
	fmt.Println("Input: s = 'aababbb'")
	fmt.Println("Output:", largestVariance("aababbb"))
	fmt.Println("Input: s = 'abcde'")
	fmt.Println("Output:", largestVariance("abcde"))
}

func largestVariance(s string) int {
	res := 0
	chars := make(map[rune]bool)
	for _, c := range s {
		chars[c] = true
	}

	for x := range chars {
		for y := range chars {
			if x == y {
				continue
			}

			totalX, totalY := 0, 0
			for _, c := range s {
				switch c {
				case x:
					totalX += 1
				case y:
					totalY += 1
				default:
					continue
				}

				if totalY > 0 {
					res = max(res, totalX-totalY)
				} else {
					res = max(res, totalX-1)
				}

				if totalX-totalY < 0 {
					totalX = 0
					totalY = 0
				}
			}
		}
	}
	return res
}
