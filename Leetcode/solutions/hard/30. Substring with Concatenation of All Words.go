package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/number-of-1-bits/
func init() {
	Solutions[30] = func() {
		fmt.Println("n = s = 'barfoothefoobarman', words = ['foo','bar']")
		fmt.Println("Output:", findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
		fmt.Println("n = s = 'wordgoodgoodgoodbestword', words = ['word','good','best','word']")
		fmt.Println("Output:", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
		fmt.Println("n = s = 'barfoofoobarthefoobarman', words = ['bar','foo','the']")
		fmt.Println("Output:", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	}
}

func findSubstring(s string, words []string) []int {
	var wordLen = len(words[0])
	var sLen = len(s)
	var results = []int{}
	var wordCount = make([]int, len(words))
	var wtoi = map[string]int{}
	for i := range words {
		wtoi[words[i]] = i
	}

	for i := range words {
		wordCount[wtoi[words[i]]]++
	}

	var sitowi = make([]int, len(s))

	for i := 0; i < sLen; i++ {
		sitowi[i] = -1
	}

	for i := 0; i < sLen-wordLen+1; i++ {
		if pos, found := wtoi[s[i:i+wordLen]]; found {
			sitowi[i] = pos
		}
	}

	wtoi = nil
	for i := 0; i < wordLen; i++ {

		var occupancy = make([]int, len(words))
		var count = 0
		var start = i
		for j := i; j < sLen-wordLen+1; j += wordLen {
			pos := sitowi[j]
			if pos != -1 && occupancy[pos] < wordCount[pos] {
				count++
				occupancy[pos]++
			} else if pos != -1 {
				for occupancy[pos] == wordCount[pos] {
					occupancy[sitowi[start]]--
					count--
					start += wordLen
				}
				count++
				occupancy[pos]++
			} else {
				occupancy = make([]int, len(words))
				count = 0
				start = j + wordLen
			}

			if count == len(words) {
				results = append(results, start)
			}
		}

	}
	return results
}
