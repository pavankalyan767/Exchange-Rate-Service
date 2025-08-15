package main

import (
	"fmt"
	"sync"
	"time"
)


type Cache struct {
	data        map[string]interface{}
	expiration  map[string]time.Time
	mutex       sync.RWMutex
	defaultTTL  time.Duration
	cleanupTick time.Duration
}

// NewCache is a constructor for the Cache struct.
// can be used for both live and historical data.
func NewCache(defaultTTL, cleanupTick time.Duration) *Cache {
	cache := &Cache{
		data:        make(map[string]interface{}), 
		expiration:  make(map[string]time.Time),
		defaultTTL:  defaultTTL,
		cleanupTick: cleanupTick,
	}

	go cache.startCleanup()

	return cache
}

// background goroutine that periodically cleanup expired items .
func (c *Cache) startCleanup() {
	ticker := time.NewTicker(c.cleanupTick)
	for {
		select {
		case <-ticker.C:
			c.cleanup()
		}
	}
}

// cleanup iterates through the cache and removes expired items.
func (c *Cache) cleanup() {
	currentTime := time.Now()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, expirationTime := range c.expiration {
		if currentTime.After(expirationTime) {
			delete(c.data, key)
			delete(c.expiration, key)
		}
	}
}

// Set adds a key-value pair to the cache with a given TTL.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	c.expiration[key] = time.Now().Add(ttl)
}

// Get retrieves a value from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, ok := c.data[key]
	if !ok || time.Now().After(c.expiration[key]) {
		return nil, false // Return nil and false if the key is not found or is expired
	}

	return value, true
}

// GetLiveRate retrieves a live rate and returns it as a float64.
func (c *Cache) GetLiveRate(key string) (float64, bool) {
	val, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	// interface to float64 type assertion
	rate, ok := val.(float64)
	if !ok {
		// Handle the case where the type is incorrect.
		fmt.Printf("Error: Value for key '%s' is not a float64.\n", key)
		return 0, false
	}
	return rate, true
}

// GetHistoryRates retrieves historical rates and returns them as a map.
func (c *Cache) GetHistoryRates(key string) (map[string]float64, bool) {
	val, ok := c.Get(key)
	if !ok {
		return nil, false
	}
	// The type assertion is now handled inside the cache package.
	rates, ok := val.(map[string]float64)
	if !ok {
		// Handle the case where the type is incorrect.
		fmt.Printf("Error: Value for key '%s' is not a map[string]float64.\n", key)
		return nil, false
	}
	return rates, true
}
