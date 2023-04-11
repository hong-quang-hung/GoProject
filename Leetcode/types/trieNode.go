package types

type TrieNode struct {
	Childrens map[byte]*TrieNode
	End       bool
}
