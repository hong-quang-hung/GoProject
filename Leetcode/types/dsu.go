package types

type DSU struct {
	Parent     []int
	Components int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DSU{
		Parent:     parent,
		Components: n,
	}
}

func (u *DSU) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

func (u *DSU) Join(x int, y int) bool {
	xFind, yFind := u.Find(x), u.Find(y)
	if xFind == yFind {
		return false
	}

	u.Components--
	return true
}
