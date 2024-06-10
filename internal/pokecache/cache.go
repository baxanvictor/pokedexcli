package pokecache

import (
	"sync"
	"time"
)

type Cache[T any] struct {
	Entries map[string]cachEntry[T]
	mu      *sync.Mutex
}

type cachEntry[T any] struct {
	createdAt *time.Time
	val       *T
}

func (c *Cache[T]) Add(key string, val *T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	c.Entries[key] = cachEntry[T]{
		createdAt: &now,
		val:       val,
	}
}

func (c *Cache[T]) Get(key string) (val *T, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.Entries[key]; ok {
		val = entry.val
		found = ok
	}
	return val, found
}

func (c *Cache[T]) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.Entries, key)
}

func NewCache[T any](duration *time.Duration, reap bool) *Cache[T] {
	cache := Cache[T]{
		Entries: make(map[string]cachEntry[T]),
		mu:      &sync.Mutex{},
	}
	if reap {
		go cache.reapLoop(duration)
	}
	return &cache
}

func (c *Cache[T]) reapLoop(duration *time.Duration) {
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

func (c *Cache[T]) removeEntriesOlderThan(duration *time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.Entries {
		if time.Since(*entry.createdAt) >= *duration {
			delete(c.Entries, key)
		}
	}
}
