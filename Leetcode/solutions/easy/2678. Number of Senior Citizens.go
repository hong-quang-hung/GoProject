package easy

import "fmt"

// Reference: https://leetcode.com/problems/number-of-senior-citizens/
func init() {
	Solutions[2678] = func() {
		fmt.Println("Input: nums = ['7868190130M7522','5303914400F9211','9273338290F4010']")
		fmt.Println("Output:", countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
		fmt.Println("Input: nums = ['1313579440F2036','2921522980M5644']")
		fmt.Println("Output:", countSeniors([]string{"1313579440F2036", "2921522980M5644"}))
	}
}

func countSeniors(details []string) int {
	res := 0
	for _, detail := range details {
		age := int(detail[11]-'0')*10 + int(detail[12]-'0')
		if age > 60 {
			res++
		}
	}
	return res
}
