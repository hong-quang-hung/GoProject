package types

type UnionFind struct {
	Parent []int
	Rank   []int
	Count  []int
}

func NewUnionFind(n int) *UnionFind {
	parent, count, rank := make([]int, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		count[i] = 1
	}

	return &UnionFind{
		Parent: parent,
		Rank:   rank,
		Count:  count,
	}
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

	count := u.Count[xset] + u.Count[yset]
	u.Count[xset] = count
	u.Count[yset] = count

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
