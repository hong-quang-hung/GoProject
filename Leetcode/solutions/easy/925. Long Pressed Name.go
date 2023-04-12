package easy

import "fmt"

// Reference: https://leetcode.com/problems/long-pressed-name/
func Leetcode_Is_Long_PressedName() {
	fmt.Println("Input: name = 'pyplrz', typed = 'ppyypllr'")
	fmt.Println("Output:", isLongPressedName("pyplrz", "ppyypllr"))
}

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	n, m := len(name), len(typed)

	for i < n && j < m {
		ch := name[i]
		if ch != typed[j] {
			return false
		}

		c1 := 0
		for i < n && name[i] == ch {
			i++
			c1++
		}

		c2 := 0
		for j < m && typed[j] == ch {
			j++
			c2++
		}

		if c1 > c2 {
			return false
		}
	}
	return j == m && i == n
}
