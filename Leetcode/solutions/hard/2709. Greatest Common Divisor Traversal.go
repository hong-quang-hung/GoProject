package hard

import "fmt"

// Reference: https://leetcode.com/problems/greatest-common-divisor-traversal/
func init() {
	Solutions[2709] = func() {
		fmt.Println(`Input: nums = [2,3,6]`)
		fmt.Println(`Output:`, canTraverseAllPairs([]int{2, 3, 6}))
		fmt.Println(`Input: nums = [3,9,5]`)
		fmt.Println(`Output:`, canTraverseAllPairs([]int{3, 9, 5}))
		fmt.Println(`Input: nums = [4,3,12,8]`)
		fmt.Println(`Output:`, canTraverseAllPairs([]int{4, 3, 12, 8}))
	}
}

func canTraverseAllPairs(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	MAX := 100000
	m := make([]bool, MAX+1)
	for _, num := range nums {
		m[num] = true
	}

	if m[1] {
		return false
	}

	sieve := make([]int, MAX+1)
	for i := 2; i <= MAX; i++ {
		if sieve[i] == 0 {
			for j := i; j <= MAX; j += i {
				sieve[j] = i
			}
		}
	}

	union := NewCanTraverseAllPairsDSU(2*MAX + 1)
	for num := range nums {
		val := num
		for val > 1 {
			prime := sieve[val]
			root := prime + MAX
			if union.Find(root) != union.Find(num) {
				union.UnionSet(root, num)
			}
			for val%prime == 0 {
				val /= prime
			}
		}
	}

	count := 0
	for i := 2; i <= MAX; i++ {
		if m[i] && union.Find(i) == i {
			count++
		}
	}
	return count == 1
}

type canTraverseAllPairsDSU struct {
	parent []int
	size   []int
}

func NewCanTraverseAllPairsDSU(n int) *canTraverseAllPairsDSU {
	parent, size := make([]int, n+1), make([]int, n+1)

	for i := 0; i <= n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &canTraverseAllPairsDSU{
		parent: parent,
		size:   size,
	}
}

func (u *canTraverseAllPairsDSU) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *canTraverseAllPairsDSU) UnionSet(x int, y int) {
	xset, yset := u.Find(x), u.Find(y)
	if xset == yset {
		return
	}

	if u.size[xset] > u.size[yset] {
		xset, yset = yset, xset
	}

	u.parent[xset] = yset
	u.size[yset] += u.size[xset]
}
