package pokecache

import (
	"fmt"
)

type Cache map[string][]byte

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Add(key string, value []byte) error {
	if key == "" {
		return fmt.Errorf("cannot add an empty key")
	}
	c[key] = value
	return nil
}

func (c Cache) Get(key string) ([]byte, bool) {
	val, ok := c[key]
	return val, ok
}
