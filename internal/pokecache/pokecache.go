package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache:    map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop()

	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cache {
			if time.Since(v.createdAt) > c.interval {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cache[key] = cacheEntry{val: val, createdAt: time.Now()}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	if entry, ok := c.cache[key]; ok {
		c.mu.Unlock()
		return entry.val, true
	}
	c.mu.Unlock()
	return nil, false
}
