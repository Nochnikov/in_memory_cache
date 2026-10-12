package cache

import (
	"sync"
	"time"
)

type item[V any] struct {
	value     V
	expiresAt time.Time
}

type Cache[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]item[V]
	once sync.Once
	stop chan struct{}
}

func (i item[V]) expired(now time.Time) bool {
	return !i.expiresAt.IsZero() && now.After(i.expiresAt)
}

func New[K comparable, V any](cleanupInterval time.Duration) *Cache[K, V] {
	c := &Cache[K, V]{
		data: make(map[K]item[V]),
		stop: make(chan struct{}),
	}

	if cleanupInterval > 0 {
		go c.janitor(cleanupInterval)
	}
	return c
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
	var exp time.Time
	if ttl > 0 {
		exp = time.Now().Add(ttl)
	}
	c.mu.Lock()
	c.data[key] = item[V]{value: value, expiresAt: exp}
	c.mu.Unlock()
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	it, ok := c.data[key]
	c.mu.RUnlock()

	if !ok {
		var zero V
		return zero, false
	}
	return it.value, true
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *Cache[K, V]) janitor(interval time.Duration) {
	ticket := time.NewTicker(interval)
	defer ticket.Stop()

	for {
		select {
		case <-ticket.C:
			c.deleteExpired()
		case <-c.stop:
			return
		}
	}
}

func (c *Cache[K, V]) deleteExpired() {
	now := time.Now()

	c.mu.Lock()
	for k, it := range c.data {
		if it.expired(now) {
			delete(c.data, k)
		}
	}
	c.mu.Lock()
}

func (c *Cache[K, V]) Close() {
	c.once.Do(func() { close(c.stop) })
}
