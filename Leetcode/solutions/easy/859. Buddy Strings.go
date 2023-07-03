package easy

import "fmt"

// Reference: https://leetcode.com/problems/buddy-strings/
func Leetcode_Buddy_Strings() {
	fmt.Println("Input: s = 'ab', goal = 'ba'")
	fmt.Println("Output:", buddyStrings("ab", "ba"))
	fmt.Println("Input: s = 'ab', goal = 'ab'")
	fmt.Println("Output:", buddyStrings("ab", "ab"))
	fmt.Println("Input: s = 'aa', goal = 'aa'")
	fmt.Println("Output:", buddyStrings("aa", "aa"))
	fmt.Println("Input: s = 'abcaa', goal = 'abcbb'")
	fmt.Println("Output:", buddyStrings("abcd", "abcbb"))
	fmt.Println("Input: s = 'abcd', goal = 'cdab'")
	fmt.Println("Output:", buddyStrings("abcd", "cdab"))
}

func buddyStrings(s string, goal string) bool {
	m1 := [26]int{}
	m2 := [26]int{}
	count := 0
	for i := 0; i < len(s); i++ {
		m1[int(s[i]-'a')]++
		m2[int(goal[i]-'a')]++
		if s[i] != goal[i] {
			count++
		}
	}
	return m1 == m2 && count == 2
}
