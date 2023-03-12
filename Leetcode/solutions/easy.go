package solutions

import (
	"strings"

	T "leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if d, v := m[target-num]; v {
			return []int{d, i}
		}
		m[num] = i
	}
	return nil
}

// Reference: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *T.TreeNode) int {
	if root == nil {
		return 0
	}
	var l int = maxDepth(root.Left)
	var r int = maxDepth(root.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

// Reference: https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	r := 0
	t := x
	for x > 0 {
		r = (r * 10) + (x % 10)
		x = x / 10

	}
	return t == r
}

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func maxProfit(prices []int) int {
	maxPrices, minPrices := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrices > prices[i] {
			minPrices = prices[i]
		} else if maxPrices < prices[i]-minPrices {
			maxPrices = prices[i] - minPrices
		}
	}
	return maxPrices
}

// Reference: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// Reference: https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	var r int = 0
	var i int = len(s) - 1
	for i >= 0 && s[i] != ' ' {
		i--
		r++
	}
	return r
}

// Reference: https://leetcode.com/problems/length-of-last-word/
func vowelStrings(words []string, left int, right int) int {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'u' || c == 'o'
	}
	var res int = 0
	for i := left; i <= right; i++ {
		w := words[i]
		l := len(w)
		if (l == 1 && isVowel(w[0])) || (isVowel(w[0]) && isVowel(w[l-1])) {
			res++
		}
	}
	return res
}
