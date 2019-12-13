package config

import (
  "go-services-base/pkg/config"
  "go-services-base/pkg/utils"
)

var Settings = struct {
  Server config.ServerSetting
  Redis  config.RedisSetting
}{}

func LoadSettings() {
  settingsInst, err := config.InitConfig("demo")
  if err != nil {
    utils.Exit("failed to load dashboard settings, err: %s", err.Error())
    return
  }
  
  if err := config.GetSettingsByKey(settingsInst, "general", &Settings.Server); err != nil {
    utils.Exit("error to read general setting, %s\n", err)
  }
  
  if err := config.GetSettingsByKey(settingsInst, "redis", &Settings.Redis); err != nil {
    utils.Exit("error to read redis setting, %s\n", err);
  }
  
  return
}
