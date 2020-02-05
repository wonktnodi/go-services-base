package main

import (
  demoConfig "github.com/wonktnodi/go-services-base/internal/config"
  "github.com/wonktnodi/go-services-base/internal/routers"
  "github.com/wonktnodi/go-services-base/pkg/cache"
  "github.com/wonktnodi/go-services-base/pkg/databases"
  
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/utils"
)

func main() {
  demoConfig.LoadSettings()
  
  logging.NewLogger(demoConfig.Settings.Server.RunMode)
  logging.SetLevel(demoConfig.Settings.Server.LogLevel)
  
  modelSettings := demoConfig.Settings.Redis
  modelSettings.DB = 1
  cache.Init(&modelSettings)
  
  databases.InitMysql(&demoConfig.Settings.Database, true, true)
  routers := routers.InitRouters()
  
  var cfg = config.ServerSetting{}
  cfg.Port = 8080
  cfg.Address = ""
  
  utils.StartService(routers, &demoConfig.Settings.Server)
}
