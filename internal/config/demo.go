package config

import (
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/mq"
  "github.com/wonktnodi/go-services-base/pkg/utils"
)

var Settings = struct {
  Server   config.ServerSetting
  Redis    config.RedisSetting
  Database config.Database
  RocketMQ mq.RocketMQ
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
    utils.Exit("error to read redis setting, %s\n", err)
  }

  if err := config.GetSettingsByKey(settingsInst, "database", &Settings.Database); err != nil {
    utils.Exit("error to read redis setting, %s\n", err)
  }
  if err := config.GetSettingsByKey(settingsInst, "rocketMQ", &Settings.RocketMQ); err != nil {
    utils.Exit("error to read rocketMQ setting, %s\n", err)
  }
  return
}
