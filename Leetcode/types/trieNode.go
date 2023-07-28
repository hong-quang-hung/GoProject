package types

type TrieNode struct {
	Childrens map[byte]*TrieNode
	End       bool
}

type TrieNodes struct {
	Suffixes [26]*TrieNodes
	Tail     bool
}

func (tr *TrieNodes) Get(ch byte) *TrieNodes {
	return tr.Suffixes[ch-'a']
}

func (tr *TrieNodes) Set(ch byte, value *TrieNodes) {
	tr.Suffixes[ch-'a'] = value
}
