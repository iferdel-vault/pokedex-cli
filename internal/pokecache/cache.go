package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		mutex: sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) error {
	if key == "" {
		return fmt.Errorf("cannot add an empty key")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.cache[key] = cacheEntry
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		currentTime := time.Now()

		c.mutex.Lock()

		for key, entry := range c.cache {
			if currentTime.Sub(entry.createdAt) > interval { // Sub method returns diff
				delete(c.cache, key)
			}
		}

		// más eff que usar defer al estár dentro de un loop y solicitarse mucho  el lock/unlock
		c.mutex.Unlock()
	}
}
