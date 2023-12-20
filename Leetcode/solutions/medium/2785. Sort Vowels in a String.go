package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/sort-vowels-in-a-string/
func init() {
	Solutions[2785] = func() {
		fmt.Println(`Input: s = "lEetcOde"`)
		fmt.Println(`Output:`, sortVowels(`lEetcOde`))
		fmt.Println(`Input: s = "lYmpH"`)
		fmt.Println(`Output:`, sortVowels(`lYmpH`))
	}
}

func sortVowels(s string) string {
	r1, r2 := []rune{}, []rune(s)
	vowels := map[rune]bool{
		'a': true, 'A': true,
		'e': true, 'E': true,
		'i': true, 'I': true,
		'o': true, 'O': true,
		'u': true, 'U': true,
	}

	for _, ch := range s {
		if vowels[ch] {
			r1 = append(r1, ch)
		}
	}

	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })

	j := 0
	for i, ch := range r2 {
		if vowels[ch] {
			r2[i] = r1[j]
			j++
		}
	}
	return string(r2)
}
