package services

import (
  "services-base/pkg/config"
  "services-base/pkg/router"
  "services-base/pkg/utils"
)

func LoadEndpoints(fileName string) *router.Endpoints {
  apiSetting, err := config.InitConfig(fileName)
  if err != nil {
    utils.Exit("failed to parse api setting, %s\n", err)
    return nil
  }
  var ApiSettings router.ServiceConfig
  err = apiSetting.Unmarshal(&ApiSettings)
  if err != nil {
    utils.Exit("failed to parse api setting, %s\n", err)
    return nil
  }
  
  endpoints := router.NewEndpoints(&ApiSettings)
  
  return endpoints
}
