package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	if key == "" {
		return fmt.Errorf("cannot add an empty key")
	}

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C { // runs every interval (i.e. interval = 5 min, runs every 5 min)
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	currentTime := time.Now().UTC()
	for key, entry := range c.cache {
		if currentTime.Sub(entry.createdAt) > interval { // Sub method returns diff
			delete(c.cache, key)
		}
	}

}
