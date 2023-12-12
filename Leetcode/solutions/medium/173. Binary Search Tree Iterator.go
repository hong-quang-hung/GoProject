package medium

import "fmt"

// Reference: https://leetcode.com/problems/binary-search-tree-iterator/
func init() {
	Solutions[173] = func() {
		fmt.Println("Input:")
		fmt.Println("[\"BSTIterator\", \"next\", \"next\", \"hasNext\", \"next\", \"hasNext\", \"next\", \"hasNext\", \"next\", \"hasNext\"]")
		fmt.Println("[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]")

		fmt.Println("Output:")
		bSTIterator := BSTIteratorConstructor(S2TreeNode("[7,3,15,null,null,9,20]"))
		fmt.Println("bSTIterator.next();", "// return", bSTIterator.Next())
		fmt.Println("bSTIterator.next();", "// return", bSTIterator.Next())
		fmt.Println("bSTIterator.hasNext();", "// return", bSTIterator.HasNext())
		fmt.Println("bSTIterator.next();", "// return", bSTIterator.Next())
		fmt.Println("bSTIterator.hasNext();", "// return", bSTIterator.HasNext())
		fmt.Println("bSTIterator.next();", "// return", bSTIterator.Next())
		fmt.Println("bSTIterator.hasNext();", "// return", bSTIterator.HasNext())
		fmt.Println("bSTIterator.next();", "// return", bSTIterator.Next())
		fmt.Println("bSTIterator.hasNext();", "// return", bSTIterator.HasNext())
	}
}

type BSTIterator struct {
	curIdx int
	values []int
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	var f func(res *[]int, root *TreeNode)

	iterator := []int{}
	f = func(arr *[]int, root *TreeNode) {
		if root == nil {
			return
		}

		f(arr, root.Left)
		*arr = append(*arr, root.Val)
		f(arr, root.Right)
	}

	f(&iterator, root)
	return BSTIterator{
		curIdx: 0,
		values: iterator,
	}
}

func (b *BSTIterator) Next() int {
	res := b.values[b.curIdx]
	b.curIdx++
	return res
}

func (b *BSTIterator) HasNext() bool {
	return b.curIdx < len(b.values)
}
