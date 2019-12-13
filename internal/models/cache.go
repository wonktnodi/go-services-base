package models

import (
  "services-base/pkg/config"
  "services-base/pkg/logging"
  "services-base/pkg/persistence"
)

var CacheStore persistence.CacheStore

func Init(setting *config.RedisSetting) (err error) {
  CacheStore, err = initCache(setting)
  if err != nil {
    logging.Errorf("failed to create cache, %s", err)
    return
  }
  
  return
}

func initCache(setting *config.RedisSetting) (ret persistence.CacheStore, err error) {
  ret, err = persistence.NewChainCache(setting)
  return
}
