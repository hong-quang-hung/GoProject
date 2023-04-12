package medium

import "fmt"

// Reference: https://leetcode.com/problems/simplify-path/
func Leetcode_Simplify_Path() {
	fmt.Println("Input: path = '/home/'")
	fmt.Println("Output:", simplifyPath("/home/"))
	fmt.Println("Input: path = '/../'")
	fmt.Println("Output:", simplifyPath("/../"))
	fmt.Println("Input: path = '/home//foo/'")
	fmt.Println("Output:", simplifyPath("/home//foo/"))
}

func simplifyPath(path string) string {
	return ""
}
