package models

import (
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/persistence"
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
