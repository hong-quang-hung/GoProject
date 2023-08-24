package hard

import (
	"fmt"
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
	return nil
}
