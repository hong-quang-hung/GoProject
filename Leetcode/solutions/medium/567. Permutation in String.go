package medium

import "fmt"

// Reference: https://leetcode.com/problems/permutation-in-string/
func Leetcode_Check_Inclusion() {
	fmt.Println("Input: s1 = 'ab', s2 = 'eidbaooo'")
	fmt.Println("Output:", checkInclusion("ab", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	s1map, s2map := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if s1map[i] == s2map[i] {
			count++
		}
	}

	for i := 0; i < n2-n1; i++ {
		r, l := s2[i+n1]-'a', s2[i]-'a'
		if count == 26 {
			return true
		}

		s2map[r]++
		if s2map[r] == s1map[r] {
			count++
		} else if s2map[r] == s1map[r]+1 {
			count--
		}

		s2map[l]--
		if s2map[l] == s1map[l] {
			count++
		} else if s2map[l] == s1map[l]-1 {
			count--
		}
	}
	return count == 26
}
