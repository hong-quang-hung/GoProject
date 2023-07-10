package easy

import "fmt"

func init() {
	Solutions[2399] = Leetcode_Check_Distances
}

// Reference: https://leetcode.com/problems/check-distances-between-same-letters/
func Leetcode_Check_Distances() {
	fmt.Println("Input: s = 'aa', distance = [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]")
	fmt.Println("Output:", checkDistances("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println("Input: s = 'abbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzza', distance = [49,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]")
	fmt.Println("Output:", checkDistances("abbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzza", []int{49, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println("Input: s = 'adaccd', distance = [1,0,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]")
	fmt.Println("Output:", checkDistances("adaccd", []int{1, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func checkDistances(s string, distance []int) bool {
	count := make([]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		idx := s[i] - 'a'
		count[idx]++
		if count[idx] == 1 {
			distance[idx] -= i
		} else {
			distance[idx] += i + 1
		}
		if count[idx] == 2 && distance[idx] != 0 {
			return false
		}
	}
	return true
}
