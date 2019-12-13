package marshaller

import (
  "github.com/vmihailenco/msgpack"
  "github.com/wonktnodi/go-services-base/pkg/persistence/cache"
  "github.com/wonktnodi/go-services-base/pkg/persistence/store"
)

// Marshaller is the struct that marshal and unmarshal cache values
type Marshaller struct {
  cache cache.Interface
}

// New creates a new marshaller that marshal/unmarshal cache values
func New(cache cache.Interface) *Marshaller {
  return &Marshaller{
    cache: cache,
  }
}

// Get obtains a value from cache and unmarshal value with given object
func (c *Marshaller) Get(key interface{}, returnObj interface{}) (interface{}, error) {
  result, err := c.cache.Get(key)
  if err != nil {
    return nil, err
  }
  
  switch result.(type) {
  case []byte:
    err = msgpack.Unmarshal(result.([]byte), returnObj)
  
  case string:
    err = msgpack.Unmarshal([]byte(result.(string)), returnObj)
  }
  
  if err != nil {
    return nil, err
  }
  
  return returnObj, nil
}

// Set sets a value in cache by marshaling value
func (c *Marshaller) Set(key, object interface{}, options *store.Options) error {
  bytes, err := msgpack.Marshal(object)
  if err != nil {
    return err
  }
  
  return c.cache.Set(key, bytes, options)
}

// Delete removes a value from the cache
func (c *Marshaller) Delete(key interface{}) error {
  return c.cache.Delete(key)
}

// Invalidate invalidate cache values using given options
func (c *Marshaller) Invalidate(options store.InvalidateOptions) error {
  return c.cache.Invalidate(options)
}

// Clear reset all cache data
func (c *Marshaller) Clear() error {
  return c.cache.Clear()
}
