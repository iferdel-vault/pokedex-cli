package pokecache

type Cache map[string][]byte

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Add(key string, value []byte) {
}

func (c Cache) Get(key string) ([]byte, bool) {
	val, ok := c[key]
	return val, ok
}
