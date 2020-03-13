package restful

import (
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/utils"
)

func LoadEndpoints(fileName string) *Endpoints {
  apiSetting, err := config.InitConfig(fileName)
  if err != nil {
    utils.Exit("failed to parse api setting, %s\n", err)
    return nil
  }
  var ApiSettings ServiceConfig
  err = apiSetting.Unmarshal(&ApiSettings)
  if err != nil {
    utils.Exit("failed to parse api setting, %s\n", err)
    return nil
  }
  
  endpoints := NewEndpoints(&ApiSettings)
  
  return endpoints
}
