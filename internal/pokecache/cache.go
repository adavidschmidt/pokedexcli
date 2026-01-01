package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cachedEntries map[string]cacheEntry
	mu            sync.Mutex
	interval      time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cachedEntries: make(map[string]cacheEntry),
		interval:      interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	currentTime := time.Now()
	c.cachedEntries[key] = cacheEntry{
		createdAt: currentTime,
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, exists := c.cachedEntries[key]
	if exists {
		return v.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cachedEntries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cachedEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
