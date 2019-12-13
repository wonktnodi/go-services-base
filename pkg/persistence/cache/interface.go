package cache

import (
  "go-services-base/pkg/persistence/codec"
  "go-services-base/pkg/persistence/store"
)

// CacheInterface represents the interface for all caches (aggregates, metric, memory, redis, ...)
type Interface interface {
  Get(key interface{}) (interface{}, error)
  Set(key, object interface{}, options *store.Options) error
  Delete(key interface{}) error
  Invalidate(options store.InvalidateOptions) error
  Clear() error
  GetType() string
}

// SetterCacheInterface represents the interface for caches that allows
// storage (for instance: memory, redis, ...)
type SetterCacheInterface interface {
  Interface
  
  GetCodec() codec.Interface
}
