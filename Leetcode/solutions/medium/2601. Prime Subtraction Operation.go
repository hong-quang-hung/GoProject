package medium

import "fmt"

// Reference: https://leetcode.com/problems/prime-subtraction-operation/
func init() {
	Solutions[2601] = func() {
		fmt.Println(`Input: nums = [4,9,6,10]`)
		fmt.Println(`Output:`, primeSubOperation([]int{4, 9, 6, 10}))
	}
}

func Leetcode_Prime_Sub_Operation() {
	fmt.Println(`Input: nums = [4,9,6,10]`)
	fmt.Println(`Output:`, primeSubOperation([]int{4, 9, 6, 10}))
}

func primeSubOperation(nums []int) bool {
	isPrime := make([]bool, 1005)
	for i := 0; i < 1005; i++ {
		isPrime[i] = true
	}

	prime := make([]int, 0)
	seive(isPrime, 1005)

	for i := 0; i < 1005; i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}

	var n int = len(nums)
	var flag bool = false
	var prev int = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < prev {
			prev = nums[i]
			continue
		}

		flag = true
		for sub := 0; sub < len(prime) && prime[sub] < nums[i]; sub++ {
			if nums[i]-prime[sub] < prev {
				prev = nums[i] - prime[sub]
				flag = false
				break
			}
		}
		if flag {
			break
		}
	}
	return !flag
}

func seive(v []bool, n int) {
	v[0] = false
	v[1] = false
	for i := 2; i*i < n; i++ {
		for j := 2 * i; j < n; j += i {
			v[j] = false
		}
	}
}
