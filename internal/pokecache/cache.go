package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) error {
	if key == "" {
		return fmt.Errorf("cannot add an empty key")
	}
	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.cache[key] = cacheEntry
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {

	entry, ok := c.cache[key]
	return entry.val, ok
}

/*
func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    for range ticker.C {
        time.Now()
    }
}
*/
