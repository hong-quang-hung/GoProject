package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/determine-if-two-strings-are-close/
func Leetcode_Close_Strings() {
	fmt.Println("Input: word1 = 'abc', word2 = 'bca'")
	fmt.Println("Output:", closeStrings("abc", "bca"))
	fmt.Println("Input: word1 = 'cabbba', word2 = 'abbccc'")
	fmt.Println("Output:", closeStrings("cabbba", "abbccc"))
	fmt.Println("Input: word1 = 'cabbb', word2 = 'abbcc'")
	fmt.Println("Output:", closeStrings("cabbb", "abbcc"))
}

func closeStrings(word1 string, word2 string) bool {
	c1, m1 := [26]int{}, [26]bool{}
	for _, ch := range word1 {
		c1[int(ch-'a')]++
		m1[int(ch-'a')] = true
	}

	c2, m2 := [26]int{}, [26]bool{}
	for _, ch := range word2 {
		c2[int(ch-'a')]++
		m2[int(ch-'a')] = true
	}

	sort.Ints(c1[:])
	sort.Ints(c2[:])
	return m1 == m2 && c1 == c2
}
