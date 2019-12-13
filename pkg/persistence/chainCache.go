package persistence

import (
  "fmt"
  "github.com/allegro/bigcache"
  "github.com/go-redis/redis/v7"
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/persistence/cache"
  "github.com/wonktnodi/go-services-base/pkg/persistence/marshaller"
  "github.com/wonktnodi/go-services-base/pkg/persistence/store"
  "time"
)

// RedisStore represents the cache with redis persistence
type ChainCache struct {
  chainCache     *cache.ChainCache
  bigCacheClient *bigcache.BigCache
  redisClient    *redis.Client
  marshal        *marshaller.Marshaller
}

const cacheTag = "chainCache"

var defaultTag = []string{cacheTag}

func NewChainCache(setting *config.RedisSetting) (val *ChainCache, err error) {
  var ret ChainCache
  
  ret.bigCacheClient, err = bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Minute))
  if err != nil {
    logging.Errorf("failed to create big cache, %s", err)
    return
  }
  ret.redisClient = redis.NewClient(&redis.Options{
    Addr: fmt.Sprintf("%s:%d", setting.Address, setting.Port),
    DB:   setting.DB,
  })
  _, err = ret.redisClient.Ping().Result()
  if err != nil {
    logging.Errorf("failed to check redis connection, %s", err)
    return
  }
  redisStore := store.NewRedis(ret.redisClient, &store.Options{Expiration: 5 * time.Second})
  memStore := store.NewBigcache(ret.bigCacheClient, nil)
  
  ret.chainCache = cache.NewChain(
    cache.New(memStore),
    cache.New(redisStore),
  )
  
  // Initializes marshal
  ret.marshal = marshaller.New(ret.chainCache)
  val = &ret
  
  return
}

func (c *ChainCache) Get(key string, value interface{}) (err error) {
  _, err = c.marshal.Get(key, value)
  return
}

// Set sets an item to the cache, replacing any existing item.
func (c *ChainCache) Set(key string, value interface{}, expire time.Duration) (err error) {
  err = c.marshal.Set(key, value, &store.Options{Expiration: expire, Tags: defaultTag})
  return
}

// Add adds an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (c *ChainCache) Add(key string, value interface{}, expire time.Duration) (err error) {
  err = c.marshal.Set(key, value, &store.Options{Expiration: expire, Tags: defaultTag})
  return
}

// Replace sets a new value for the cache key only if it already exists. Returns an
// error if it does not.
func (c *ChainCache) Replace(key string, data interface{}, expire time.Duration) (err error) {
  err = c.marshal.Set(key, data, &store.Options{Expiration: expire, Tags: defaultTag})
  return
}

// Delete removes an item from the cache. Does nothing if the key is not in the cache.
func (c *ChainCache) Delete(key string) (err error) {
  err = c.marshal.Delete(key)
  return
}

// Increment increments a real number, and returns error if the value is not real
func (*ChainCache) Increment(key string, data uint64) (ret uint64, err error) {
  return
}

// Decrement decrements a real number, and returns error if the value is not real
func (*ChainCache) Decrement(key string, data uint64) (ret uint64, err error) {
  return
}

// Flush deletes all items from the cache.
func (c *ChainCache) Flush() (err error) {
  //c.redisClient.FlushAll()
  c.bigCacheClient.Reset()
  return
}
