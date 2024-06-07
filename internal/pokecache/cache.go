package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cachEntry
	mu      *sync.Mutex
}

type cachEntry struct {
	createdAt *time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	c.Entries[key] = cachEntry{
		createdAt: &now,
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.Entries[key]; ok {
		val = entry.val
		found = ok
	}
	return val, found
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.Entries, key)
}

func NewCache(duration *time.Duration) *Cache {
	cache := Cache{
		Entries: make(map[string]cachEntry),
		mu:      &sync.Mutex{},
	}
	go cache.reapLoop(duration)
	return &cache
}

func (c *Cache) reapLoop(duration *time.Duration) {
	dur := duration
	if dur == nil || duration.Seconds() <= 0 {
		defaultDuration := time.Second * 10
		dur = &defaultDuration
	}
	ticker := time.NewTicker(*dur)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.removeEntriesOlderThan(dur)
	}
}

func (c *Cache) removeEntriesOlderThan(duration *time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.Entries {
		if time.Since(*entry.createdAt) >= *duration {
			delete(c.Entries, key)
		}
	}
}
