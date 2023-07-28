package medium

import "fmt"

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
	trie := new(TrieNode)
	for _, product := range products {
		node := trie
		for i := range product {
			ch := product[i]
			if node.Childrens == nil {
				node.Childrens = make(map[byte]*TrieNode)
			}
			if node.Childrens[ch] == nil {
				node.Childrens[ch] = new(TrieNode)
			}
			node = node.Childrens[ch]
		}
		node.End = true
	}

	res := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		search := []string{}
		node := trie
		prefix := searchWord[:i]
		for j := range prefix {
			node = node.Childrens[prefix[j]]
			if node == nil {
				break
			}

			search = append(search, string(prefix[j]))
			if len(search) == 3 {
				break
			}
		}

		res = append(res, search)
	}
	return res
}
