package types

type UnionFind struct {
	Parent []int
	Rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	rank := make([]int, size)
	return &UnionFind{Parent: parent, Rank: rank}
}

func (u *UnionFind) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

func (u *UnionFind) UnionSet(x int, y int) bool {
	xset, yset := u.Find(x), u.Find(y)
	if xset == yset {
		return false
	}

	if u.Rank[xset] < u.Rank[yset] {
		u.Parent[xset] = yset
	} else if u.Rank[xset] > u.Rank[yset] {
		u.Parent[yset] = xset
	} else {
		u.Parent[yset] = xset
		u.Rank[xset]++
	}
	return true
}
