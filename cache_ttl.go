package cache

import (
	"errors"
	"sync"
	"time"
)

type KeyInfo struct {
	key interface{}
	cleanTime int64
}

type Cache struct {
	currentCache map[string]KeyInfo
	ttl time.Duration
	mu sync.RWMutex
}

func (m *Cache) findForClean() {
	for key, value := range m.currentCache {
		if time.Now().Unix() > value.cleanTime {
			delete(m.currentCache, key)
		}
	}
}

func (m *Cache) scanCache() {
	go func() {
		for {
			m.findForClean()
		}
	}()
}

func New() *Cache {
	cache := &Cache{
		currentCache: make(map[string]KeyInfo),
	}
	cache.scanCache()
	return cache
}

func (m *Cache) Set(key string, value interface{}, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	cleanTime := time.Now().Add(ttl).Unix()
	m.currentCache[key] = KeyInfo{
		key: value,
		cleanTime: cleanTime,
	}
}

func (m *Cache) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.currentCache, key)
}

func (m *Cache) Get(key string) (interface{}, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, exists := m.currentCache[key]
	if exists {
		return value.key, nil
	}
	return nil, errors.New("key not found")
}
