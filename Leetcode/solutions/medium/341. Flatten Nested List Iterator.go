package medium

import "fmt"

// Reference: https://leetcode.com/problems/flatten-nested-list-iterator/
func init() {
	Solutions[341] = func() {
		fmt.Println("Input: nestedList = [[1,1],2,[1,1]]")
		fmt.Println("Input: nestedList = [1,[4,[6]]]")
	}
}

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool           { return true }
func (n NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)     {}
func (n *NestedInteger) Add(elem NestedInteger)   {}
func (n NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
	list []int
	idx  int
}

func NestedIteratorConstructor(nestedList []*NestedInteger) *NestedIterator {
	flattenedList := []int{}
	var flatten func(*[]int, []*NestedInteger)
	flatten = func(flattenedList *[]int, nestedList []*NestedInteger) {
		for _, ni := range nestedList {
			if ni.IsInteger() {
				*flattenedList = append(*flattenedList, ni.GetInteger())
				continue
			}
			flatten(flattenedList, ni.GetList())
		}
	}

	flatten(&flattenedList, nestedList)
	return &NestedIterator{list: flattenedList}
}

func (n *NestedIterator) Next() int {
	if !n.HasNext() {
		return -1
	}

	val := n.list[n.idx]
	n.idx++
	return val
}

func (n *NestedIterator) HasNext() bool {
	return n.idx < len(n.list)
}
