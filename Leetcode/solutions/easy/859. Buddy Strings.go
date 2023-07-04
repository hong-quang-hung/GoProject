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
}

func buddyStrings(s string, goal string) bool {
	n := len(s)
	if n != len(goal) {
		return false
	}

	hash := make(map[byte]int)
	count := 0
	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			count++
			if count > 2 {
				return false
			}
		}

		hash[s[i]]++
		hash[goal[i]]--
	}

	if s == goal && len(hash) == n {
		return false
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
