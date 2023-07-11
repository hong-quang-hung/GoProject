package easy

import "fmt"

// Reference: https://leetcode.com/problems/greatest-common-divisor-of-strings/
func init() {
	Solutions[1071] = func() {
		fmt.Println("Input: str1 = 'ABCABC', str2 = 'ABC")
		fmt.Println("Output:", gcdOfStrings("ABCABC", "ABC"))
	}
}

func gcdOfStrings(str1 string, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	if l1 == l2 {
		if str1 == str2 {
			return str1
		} else {
			return ""
		}
	} else if l1 > l2 {
		f := str1[0 : l1-l2]
		l := str1[l2:l1]
		if f != l {
			return ""
		}
		return gcdOfStrings(l, str2)
	} else {
		f := str2[0 : l2-l1]
		l := str2[l1:l2]
		if f != l {
			return ""
		}
		return gcdOfStrings(str1, l)
	}
}
