package goutils

type CountList struct {
	internalMap map[int64]int
}

func NewCountList() *CountList {
	return &CountList{
		internalMap: make(map[int64]int),
	}
}

func (c *CountList) Add(item Hashable) {
	hash := item.Hash()
	_, ok := c.internalMap[hash]
	if ok {
		c.internalMap[hash]++
	} else {
		c.internalMap[hash] = 1
	}
}

func (c *CountList) Get(item Hashable) int {
	ret, ok := c.internalMap[item.Hash()]
	if !ok {
		return 0
	}
	return ret
}

func (c *CountList) Contains(item Hashable) bool {
	_, ok := c.internalMap[item.Hash()]
	return ok
}

func (c *CountList) Clear() {
	c.internalMap = make(map[int64]int)
}

func (c *CountList) Size() int {
	return len(c.internalMap)
}

func (c *CountList) IsEmpty() bool {
	return c.Size() == 0
}

func (c *CountList) Keys() []int64 {
	keys := make([]int64, 0, len(c.internalMap))
	for k := range c.internalMap {
		keys = append(keys, k)
	}
	return keys
}

func (c *CountList) Values() []int {
	values := make([]int, 0, len(c.internalMap))
	for _, v := range c.internalMap {
		values = append(values, v)
	}
	return values
}

func (c *CountList) Remove(item Hashable) {
	delete(c.internalMap, item.Hash())
}
