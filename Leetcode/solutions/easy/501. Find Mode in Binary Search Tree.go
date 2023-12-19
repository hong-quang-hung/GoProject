package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-mode-in-binary-search-tree/
func init() {
	Solutions[501] = func() {
		fmt.Println(`Input: root = [1,null,2,2]`)
		fmt.Println(`Output:`, findMode(S2TreeNode(`[1,null,2,2]`)))
		fmt.Println(`Input: root = [0]`)
		fmt.Println(`Output:`, findMode(S2TreeNode(`[0]`)))
	}
}

func findModeInorder(root *TreeNode, maxFreq, currVal, freq *int, sol *[]int) {
	if root.Left != nil {
		findModeInorder(root.Left, maxFreq, currVal, freq, sol)
	}

	if root.Val == (*currVal) || (*freq) == 0 {
		(*freq)++
	} else if root.Val != (*currVal) {
		(*freq) = 1
	}

	(*currVal) = root.Val
	if (*freq) > (*maxFreq) {
		(*maxFreq) = (*freq)
		(*sol) = []int{root.Val}
	} else if (*freq) == (*maxFreq) {
		(*sol) = append((*sol), root.Val)
	}

	if root.Right != nil {
		findModeInorder(root.Right, maxFreq, currVal, freq, sol)
	}
}

func findMode(root *TreeNode) []int {
	maxFreq := 0
	freq := 0
	currVal := 0
	res := []int{}
	findModeInorder(root, &maxFreq, &currVal, &freq, &res)
	return res
}
