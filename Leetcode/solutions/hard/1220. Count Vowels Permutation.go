package hard

import "fmt"

// Reference: https://leetcode.com/problems/count-vowels-permutation/
func init() {
	Solutions[1220] = func() {
		fmt.Println(`Input: n = 1`)
		fmt.Println(`Output:`, countVowelPermutation(1))
		fmt.Println(`Input: n = 2`)
		fmt.Println(`Output:`, countVowelPermutation(2))
		fmt.Println(`Input: n = 5`)
		fmt.Println(`Output:`, countVowelPermutation(5))
	}
}

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 2; j <= n; j++ {
		n_a := e
		n_e := (a + i) % mod
		n_i := (a + e + o + u) % mod
		n_o := (i + u) % mod
		n_u := a
		a, e, i, o, u = n_a, n_e, n_i, n_o, n_u
	}
	return (a + e + i + o + u) % mod
}
