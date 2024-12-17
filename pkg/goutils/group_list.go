package goutils

type GroupList struct {
	internalMap map[int64][]Hashable
}

func NewGroupList() *GroupList {
	return &GroupList{
		internalMap: make(map[int64][]Hashable),
	}
}

func (g *GroupList) Add(item Hashable) {
	hash := item.Hash()
	if _, ok := g.internalMap[hash]; !ok {
		g.internalMap[hash] = make([]Hashable, 0)
	}
	g.internalMap[hash] = append(g.internalMap[hash], item)
}

func (g *GroupList) Get(item Hashable) []Hashable {
	hash := item.Hash()
	if _, ok := g.internalMap[hash]; !ok {
		g.internalMap[hash] = make([]Hashable, 0)
	}
	return g.internalMap[hash]
}

func (g *GroupList) Length() int {
	return len(g.internalMap)
}

func (g *GroupList) Clear() {
	g.internalMap = make(map[int64][]Hashable)
}

func (g *GroupList) Keys() []int64 {
	keys := make([]int64, 0)
	for k, _ := range g.internalMap {
		keys = append(keys, k)
	}
	return keys
}

func (g *GroupList) Values() [][]Hashable {
	values := make([][]Hashable, 0)
	for _, v := range g.internalMap {
		values = append(values, v)
	}
	return values
}
