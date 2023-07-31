package types

import "sort"

type TrieNode struct {
	Childrens map[byte]*TrieNode
	End       bool
}

type TrieNodes struct {
	Suffixes  [26]*TrieNodes
	Character []byte
	Tail      bool
}

func (tr *TrieNodes) Get(ch byte) *TrieNodes {
	return tr.Suffixes[ch-'a']
}

func (tr *TrieNodes) Set(ch byte, value *TrieNodes) {
	tr.Suffixes[ch-'a'] = value
	idx := sort.Search(len(tr.Character), func(i int) bool { return tr.Character[i] >= ch })
	if len(tr.Character) == idx {
		tr.Character = append(tr.Character, ch)
	} else if tr.Character[idx] != ch {
		tr.Character = append(tr.Character[:idx+1], tr.Character[idx:]...)
		tr.Character[idx] = ch
	}
}
