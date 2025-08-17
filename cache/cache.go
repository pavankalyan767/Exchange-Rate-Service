package cache

import (
	"sync"
	"time"

	"github.com/go-kit/log"
)

type Cache struct {
	data        map[string]interface{}
	expiration  map[string]time.Time
	mutex       sync.RWMutex
	defaultTTL  time.Duration
	cleanupTick time.Duration
	logger      log.Logger
}

// NewCache is a constructor for the Cache struct.
// can be used for both live and historical data.
func NewCache(defaultTTL, cleanupTick time.Duration, logger log.Logger) *Cache {
	cache := &Cache{
		data:        make(map[string]interface{}),
		expiration:  make(map[string]time.Time),
		defaultTTL:  defaultTTL,
		cleanupTick: cleanupTick,
		logger:      logger,
	}

	go cache.startCleanup()

	return cache
}

// background goroutine that periodically cleansup expired items .
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
func (c *Cache) GetRateWithDate(date string, currencyPair string) (float64, bool) {
	// First, retrieve the entire map of rates for the given date.
	val, ok := c.Get(date)
	if !ok {
		// No data found for this date.
		return 0, false
	}

	// Now, perform a type assertion to get the map of rates.
	ratesMap, ok := val.(map[string]float64)
	if !ok {
		// The cached value is not a map[string]float64. This indicates a data integrity issue.
		c.logger.Log("Error: Value for date '%s' is not a map[string]float64.\n", date)
		return 0, false
	}

	// Finally, get the specific currency pair from the map.
	rate, ok := ratesMap[currencyPair]
	c.logger.Log("Retrieved rate:", rate, "for currency pair:", currencyPair)
	return rate, ok
}

// GetHistoryRates retrieves historical rates and returns them as a map.
