package cache

import (
  "crypto"
  "fmt"
  "reflect"
  "go-services-base/pkg/persistence/codec"
  "go-services-base/pkg/persistence/store"
  "strings"
)

const (
  // CacheType represents the cache type as a string value
  CachingType = "cache"
)

// Cache represents the configuration needed by a cache
type Cache struct {
  codec codec.Interface
}

// New instantiates a new cache entry
func New(store store.Interface) *Cache {
  return &Cache{
    codec: codec.New(store),
  }
}

// Get returns the object stored in cache if it exists
func (c *Cache) Get(key interface{}) (interface{}, error) {
  cacheKey := c.getCacheKey(key)
  return c.codec.Get(cacheKey)
}

// Set populates the cache item using the given key
func (c *Cache) Set(key, object interface{}, options *store.Options) error {
  cacheKey := c.getCacheKey(key)
  return c.codec.Set(cacheKey, object, options)
}

// Delete removes the cache item using the given key
func (c *Cache) Delete(key interface{}) error {
  cacheKey := c.getCacheKey(key)
  return c.codec.Delete(cacheKey)
}

// Invalidate invalidates cache item from given options
func (c *Cache) Invalidate(options store.InvalidateOptions) error {
  return c.codec.Invalidate(options)
}

// Clear resets all cache data
func (c *Cache) Clear() error {
  return c.codec.Clear()
}

// GetCodec returns the current codec
func (c *Cache) GetCodec() codec.Interface {
  return c.codec
}

// GetType returns the cache type
func (c *Cache) GetType() string {
  return CachingType
}

// getCacheKey returns the cache key for the given key object by computing a
// checksum of key struct
func (c *Cache) getCacheKey(key interface{}) string {
  return strings.ToLower(checksum(key))
}

// checksum hashes a given object into a string
func checksum(object interface{}) string {
  digested := crypto.MD5.New()
  fmt.Fprint(digested, reflect.TypeOf(object))
  fmt.Fprint(digested, object)
  hash := digested.Sum(nil)
  
  return fmt.Sprintf("%x", hash)
}
