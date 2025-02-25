package main

import (
	"fmt"
	"sync"
)

// Создайте простой кэш, который будет хранить значения по ключу.
// Реализуйте методы для получения и установки значений

type CacheMethods interface {
	Set(key, val string)
	Get(key string) (string, bool)
}

type Cache struct {
	m map[string]string
	mu sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]string),
		//mu: sync.Mutex,
	}
}

func (c *Cache) Set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = val
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.m[key]
	return val, ok
}

type CacheSyncMap struct {
	m sync.Map 
}

func NewSyncCache() *CacheSyncMap {
	return &CacheSyncMap{}
}

func (c *CacheSyncMap) Set(key, val string) {
	c.m.Store(key, val)
}

func (c *CacheSyncMap) Get(key string) (string, bool) {
	value, exists := c.m.Load(key)
	if !exists {
		return "", false
	}
	return value.(string), true
}

func map_cache_test() {
	cache := NewCache()
	//cache := NewSyncCache()
	var wg sync.WaitGroup

	// Заполнение кэша
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		}(i)
	}

	wg.Wait()

	// Чтение из кэша
	for i := 0; i < 10; i++ {
		if value, exists := cache.Get(fmt.Sprintf("key%d", i)); exists {
			fmt.Printf("Found: key%d = %s\n", i, value)
		} else {
			fmt.Printf("key%d not found\n", i)
		}
	}
}