package solutions

import (
	"strings"

	"leetcode.com/Leetcode/types"
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

// Reference: https://leetcode.com/problems/unique-number-of-occurrences/
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}
	s := make(map[int]bool)
	for _, v := range m {
		if _, ok := s[v]; ok {
			return false
		} else {
			s[v] = true
		}
	}
	return true
}

// Reference: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *types.TreeNode) int {
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

// Reference: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			var j int = 1
			for ; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}

// Reference: https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	var res int = 0
	for _, num := range nums {
		if num != val {
			nums[res] = num
			res++
		}
	}
	return res
}

// Reference: https://leetcode.com/problems/number-of-even-and-odd-bits/
func evenOddBit(n int) []int {
	arr := []int{0, 0}
	var p int = 0
	for n > 0 {
		if n%2 == 1 {
			arr[p%2]++
		}
		p++
		n /= 2
	}
	return arr
}

// Reference: https://leetcode.com/problems/k-items-with-the-maximum-sum/
func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if numOnes >= k {
		return k
	}
	if numZeros >= k-numOnes {
		return numOnes
	}
	return 2*numOnes - k + numZeros
}

// Reference: https://leetcode.com/problems/verifying-an-alien-dictionary/
func isAlienSorted(words []string, order string) bool {
	hash := make(map[byte]int)
	for i := 0; i < 26; i++ {
		hash[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}

			if words[i][j] != words[i+1][j] {
				if hash[words[i][j]] > hash[words[i+1][j]] {
					return false
				} else {
					break
				}
			}
		}
	}
	return true
}

// Reference: https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	var stack = []rune{}
	for _, c := range s {
		if c == '(' {
			stack = append([]rune{')'}, stack...)
		} else if c == '[' {
			stack = append([]rune{']'}, stack...)
		} else if c == '{' {
			stack = append([]rune{'}'}, stack...)
		} else {
			if len(stack) == 0 {
				return false
			}

			last := stack[0]
			stack = stack[1:]
			if last != c {
				return false
			}
		}
	}
	return len(stack) == 0
}

// Reference: https://leetcode.com/problems/greatest-common-divisor-of-strings/
func gcdOfStrings(str1 string, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	if l1 == l2 {
		if str1 == str2 {
			return str1
		} else {
			return ""
		}
	} else if l1 > l2 {
		f := str1[0 : l1-l2]
		l := str1[l2:l1]
		if f != l {
			return ""
		}
		return gcdOfStrings(l, str2)
	} else {
		f := str2[0 : l2-l1]
		l := str2[l1:l2]
		if f != l {
			return ""
		}
		return gcdOfStrings(str1, l)
	}
}

// Reference: https://leetcode.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	distinct := make(map[int]int)
	for i, num := range nums {
		if d, c := distinct[num]; c && i-d <= k {
			return true
		} else {
			distinct[num] = i
		}
	}
	return false
}

// Reference: https://leetcode.com/problems/binary-search/
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
