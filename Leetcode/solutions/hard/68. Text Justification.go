package hard

import (
	"fmt"
	"math"
	"strings"
)

// Reference: https://leetcode.com/problems/text-justification/
func init() {
	Solutions[68] = func() {
		fmt.Println("Input: words = ['This', 'is', 'an', 'example', 'of', 'text', 'justification.'], maxWidth = 16")
		fmt.Println("Output:", fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
		fmt.Println("Input: words = ['What','must','be','acknowledgment','shall','be'], maxWidth = 16")
		fmt.Println("Output:", fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
		fmt.Println("Input: words = ['Science','is','what','we','understand','well','enough','to','explain','to','a','computer.','Art','is','everything','else','we','do'], maxWidth = 20")
		fmt.Println("Output:", fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
	}
}

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	for i := 0; i < len(words); i++ {
		cache := []string{}
		size := 0
		j := i
		for ; j < len(words); j++ {
			if size+len(words[j]) <= maxWidth {
				cache = append(cache, words[j])
				size += len(words[j]) + 1
			} else {
				j--
				break
			}
		}
		newLine := buildString(cache, maxWidth, j == len(words))
		i = j
		res = append(res, newLine)
	}
	return res
}

func buildString(words []string, maxWidth int, isLast bool) string {
	res := ""
	tempResult := strings.Join(words, " ")
	if isLast || len(words) == 1 {
		res = tempResult
		for i := len(res); i < maxWidth; i++ {
			res += " "
		}
		return res
	}

	totalSpace := maxWidth - len(tempResult) + len(words) - 1
	for i := 0; i < len(words); i++ {
		res += words[i]
		numSpaces := int(math.Ceil(float64(totalSpace) / float64(len(words)-1-i)))
		totalSpace -= numSpaces
		for s := 0; s < numSpaces; s++ {
			res += " "
		}
	}
	return res
}
