package medium

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/search-suggestions-system/
func init() {
	Solutions[1268] = func() {
		fmt.Println("Input: products = ['mobile','mouse','moneypot','monitor','mousepad'], searchWord = 'mouse'")
		fmt.Println("Output:", suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
		fmt.Println("Input: products = ['havana'], searchWord = 'havana")
		fmt.Println("Output:", suggestedProducts([]string{"havana"}, "havana"))
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := new(TrieNodes)
	for _, product := range products {
		current := trie
		for i := range product {
			if current.Get(product[i]) == nil {
				current.Set(product[i], new(TrieNodes))
			}
			current = current.Get(product[i])
		}
		current.Tail = true
	}

	res := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		current := trie
		for j := range prefix {
			current = current.Get(prefix[j])
			if current == nil {
				break
			}
		}

		search := []string{}
		if current.Tail {
			search = append(search, prefix)
		}

		sb := new(strings.Builder)
		sb.WriteString(prefix)
		stack := []*TrieNodes{current}
		for len(stack) > 0 {
			stack = stack[1:]
		}

		res = append(res, search)
	}
	return res
}

func suggestedProductsDFS() {

}
