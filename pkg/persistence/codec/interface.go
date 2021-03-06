package codec

import (
  "github.com/wonktnodi/go-services-base/pkg/persistence/store"
)

// CodecInterface represents an instance of a cache codec
type Interface interface {
  Get(key interface{}) (interface{}, error)
  Set(key interface{}, value interface{}, options *store.Options) error
  Delete(key interface{}) error
  Invalidate(options store.InvalidateOptions) error
  Clear() error
  
  GetStore() store.Interface
  GetStats() *Stats
}
