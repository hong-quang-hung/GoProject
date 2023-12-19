package medium

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/simplify-path/
func init() {
	Solutions[71] = func() {
		fmt.Println(`Input: path = '/.'`)
		fmt.Println(`Output:`, simplifyPath(`/.`))
		fmt.Println(`Input: path = '/a/../../b/../c//.//'`)
		fmt.Println(`Output:`, simplifyPath(`/a/../../b/../c//.//`))
		fmt.Println(`Input: path = '/a//b////c/d//././/..'`)
		fmt.Println(`Output:`, simplifyPath(`/a//b////c/d//././/..`))
	}
}

func simplifyPath(path string) string {
	contents := strings.Split(path, `/`)
	lenContents := len(contents)
	for i := 0; i < lenContents; i++ {
		char := contents[i]
		if (char == `..` && i == 0) || (char == `` && i < lenContents) || char == `.` {
			contents = append(contents[:i], contents[i+1:]...)
			i--
			lenContents--
		} else if char == `..` {
			contents = append(contents[:i-1], contents[i+1:]...)
			i = i - 2
			lenContents = lenContents - 2
		}
	}
	return `/` + strings.Join(contents, `/`)
}
