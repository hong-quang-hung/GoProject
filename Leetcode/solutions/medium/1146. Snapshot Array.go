package medium

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[1146] = Leetcode_Snapshot_Array
}

// Reference: https://leetcode.com/problems/snapshot-array/
func Leetcode_Snapshot_Array() {
	fmt.Println("Input:")
	fmt.Println("['SnapshotArray','set','snap','set','get']")
	fmt.Println("Output:")

	snapshotArr := SnapshotArrayConstructor(3)
	snapshotArr.Set(0, 5)
	fmt.Println("snapshotArr.snap()", "=>", snapshotArr.Snap())
	snapshotArr.Set(0, 6)
	fmt.Println("snapshotArr.get(0, 0)", "=>", snapshotArr.Get(0, 0))
}

type SnapshotArray struct {
	arr     [][][2]int
	snap_id int
}

func SnapshotArrayConstructor(length int) SnapshotArray {
	snap := new(SnapshotArray)
	snap.arr = make([][][2]int, length)
	for i := 0; i < length; i++ {
		snap.arr[i] = make([][2]int, 1)
		snap.arr[i][0] = [2]int{0, 0}
	}
	return *snap
}

func (s *SnapshotArray) Set(index int, val int) {
	h := s.arr[index]
	if h[len(h)-1][0] < s.snap_id && h[len(h)-1][1] != val {
		s.arr[index] = append(h, [2]int{s.snap_id, val})
	} else {
		s.arr[index][len(h)-1][1] = val
	}
}

func (s *SnapshotArray) Snap() int {
	s.snap_id++
	return s.snap_id - 1
}

func (s *SnapshotArray) Get(index int, snap_id int) int {
	h := s.arr[index]
	i := sort.Search(len(h), func(j int) bool {
		return (j < len(h)-1 && h[j+1][0] > snap_id) || j == len(h)-1
	})
	return h[i][1]
}
