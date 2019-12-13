package main

import (
  "fmt"
  demoConfig "services-base/internal/config"
  "services-base/internal/routers"
  
  "services-base/internal/models"
  "services-base/pkg/config"
  "services-base/pkg/logging"
  "services-base/pkg/utils"
)

func main() {
  demoConfig.LoadSettings()
  
  logging.NewLogger(demoConfig.Settings.Server.RunMode)
  logging.SetLevel(demoConfig.Settings.Server.LogLevel)
  
  modelSettings := demoConfig.Settings.Redis
  modelSettings.DB = 1
  models.Init(&modelSettings)
  
  routers := routers.InitRouters()
  
  fmt.Print("daadd")
  var cfg = config.ServerSetting{}
  cfg.Port = 8080
  cfg.Address = ""
  
  utils.StartService(routers, &demoConfig.Settings.Server)
}
