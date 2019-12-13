package main

import (
  "fmt"
  demoConfig "github.com/wonktnodi/go-services-base/internal/config"
  "github.com/wonktnodi/go-services-base/internal/routers"
  
  "github.com/wonktnodi/go-services-base/internal/models"
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
  models.Init(&modelSettings)
  
  routers := routers.InitRouters()
  
  fmt.Print("daadd")
  var cfg = config.ServerSetting{}
  cfg.Port = 8080
  cfg.Address = ""
  
  utils.StartService(routers, &demoConfig.Settings.Server)
}
