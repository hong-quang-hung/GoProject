package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/search-suggestions-system/
func init() {
	Solutions[1268] = func() {
		fmt.Println("Input: products = ['mobile','mouse','moneypot','monitor','mousepad'], searchWord = 'mouse'")
		fmt.Println("Output:", suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
		fmt.Println("Input: products = ['havana'], searchWord = 'havana'")
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
	current := trie
	str := []byte{}
	for i := 0; i < len(searchWord); i++ {
		if current != nil {
			current = current.Get(searchWord[i])
			str = append(str, searchWord[i])
		}
		res = append(res, suggestedProductsGetList(current, str))
	}
	return res
}

func suggestedProductsGetList(current *TrieNodes, str []byte) []string {
	res := []string{}
	if current != nil {
		suggestedProductsGetListDFS(current, &res, str)
	}
	return res
}

func suggestedProductsGetListDFS(current *TrieNodes, res *[]string, str []byte) {
	if len(*res) == 3 {
		return
	}

	if current.Tail {
		*res = append(*res, string(str))
	}

	for _, ch := range current.Character {
		str = append(str, ch)
		suggestedProductsGetListDFS(current.Get(ch), res, str)
		str = str[:len(str)-1]
	}
}
