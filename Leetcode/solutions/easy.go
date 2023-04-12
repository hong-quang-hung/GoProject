package solutions

import (
	"math"
	"sort"
	"strconv"
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

// Reference: https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/
func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				return nums1[i]
			}
		}
	}
	if nums1[0] < nums2[0] {
		return nums1[0]*10 + nums2[0]
	}
	return nums2[0]*10 + nums1[0]
}

// Reference: https://leetcode.com/problems/find-the-longest-balanced-substring-of-a-binary-string/
func findTheLongestBalancedSubstring(s string) int {
	res := 0
	temp := "01"
	for len(temp) <= len(s) {
		if strings.Contains(s, temp) {
			res = len(temp)
		}
		temp = "0" + temp + "1"
	}
	return res
}

// Reference: https://leetcode.com/problems/shuffle-the-array/
func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[i+len(nums)/2])
	}
	return res
}

// Reference: https://leetcode.com/problems/long-pressed-name/
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	n, m := len(name), len(typed)

	for i < n && j < m {
		ch := name[i]
		if ch != typed[j] {
			return false
		}

		c1 := 0
		for i < n && name[i] == ch {
			i++
			c1++
		}

		c2 := 0
		for j < m && typed[j] == ch {
			j++
			c2++
		}

		if c1 > c2 {
			return false
		}
	}
	return j == m && i == n
}

// Reference: https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	s, c := "", false
	for len(a) > 0 || len(b) > 0 {
		bitA := len(a) > 0 && a[len(a)-1] == '1'
		bitB := len(b) > 0 && b[len(b)-1] == '1'

		if xor := c != bitA; xor != bitB {
			s = "1" + s
		} else {
			s = "0" + s
		}
		c = (bitA && bitB) || (c && (bitA || bitB))
		if len(a) > 0 {
			a = a[0 : len(a)-1]
		}
		if len(b) > 0 {
			b = b[0 : len(b)-1]
		}
	}
	if c {
		return "1" + s
	}
	return s
}

// Reference: https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	mid, left, right := 0, 0, len(nums)-1
	for left <= right {
		mid = int(math.Floor(float64(left+right) / float64(2)))
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// Reference: https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {
	head := new(types.ListNode)
	newList := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newList.Next = &types.ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			newList.Next = &types.ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		newList = newList.Next
	}

	for list1 != nil {
		newList.Next = &types.ListNode{Val: list1.Val}
		list1 = list1.Next
		newList = newList.Next
	}

	for list2 != nil {
		newList.Next = &types.ListNode{Val: list2.Val}
		list2 = list2.Next
		newList = newList.Next
	}
	return head.Next
}

// Reference: https://leetcode.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]bool)
	for i := range pattern {
		if v, c := m1[pattern[i]]; c {
			if v != arr[i] {
				return false
			}
		} else {
			if m2[arr[i]] {
				return false
			}

			m1[pattern[i]] = arr[i]
			m2[arr[i]] = true
		}
	}
	return true
}

// Reference: https://leetcode.com/problems/split-with-minimum-sum/
func splitNum(num int) int {
	arr := make([]int, len(strconv.Itoa(num)))
	i, numA, numB := 0, 0, 0
	for num > 0 {
		arr[i] = num % 10
		num = num / 10
		i++
	}

	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			numA = numA*10 + arr[i]
		} else {
			numB = numB*10 + arr[i]
		}
	}
	return numA + numB
}

// Reference: https://leetcode.com/problems/left-and-right-sum-differences/
func leftRigthDifference(nums []int) []int {
	leftSum, rightSum := 0, 0
	for _, num := range nums {
		rightSum += num
	}

	ans := make([]int, len(nums))
	for i, num := range nums {
		rightSum -= num
		ans[i] = int(math.Abs(float64(leftSum - rightSum)))
		leftSum += num
	}
	return ans
}

// Reference: https://leetcode.com/problems/invert-binary-tree/
func invertTree(root *types.TreeNode) *types.TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

// Reference: https://leetcode.com/problems/minimum-distance-between-bst-nodes
func minDiffInBST(root *types.TreeNode) int {
	minDiff, preValue := math.MaxInt, -1
	inorderTraversal(root, &preValue, &minDiff)
	return minDiff
}

func inorderTraversal(root *types.TreeNode, preValue *int, minDiff *int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, preValue, minDiff)

	if *preValue != -1 && root.Val-*preValue < *minDiff {
		*minDiff = root.Val - *preValue
	}

	*preValue = root.Val
	inorderTraversal(root.Right, preValue, minDiff)
}

// Reference: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func removeAdjacentDuplicates(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

// Reference: https://leetcode.com/problems/can-place-flowers/
func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	for i := 0; i < size; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == size-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			i++
		}
	}
	return n <= 0
}

// Reference: https://leetcode.com/problems/add-to-array-form-of-integer/
func addToArrayForm(num []int, k int) []int {
	var ext int
	n := len(num)
	i := n - 1
	for k > 0 {
		if i >= 0 {
			ext = (num[i] + k%10) / 10
			num[i] = (num[i] + k%10) - ext*10
		} else {
			ext = 0
			num = append([]int{k % 10}, num...)
		}
		k = k/10 + ext
		i--
	}
	return num
}

// Reference: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int][]int)
	for _, num := range nums1 {
		m[num[0]] = num
	}

	for _, num := range nums2 {
		if m[num[0]] == nil {
			m[num[0]] = []int{num[0], num[1]}
		} else {
			m[num[0]][1] += num[1]
		}
	}

	res := make([][]int, 0)
	for _, v := range m {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	return res
}
