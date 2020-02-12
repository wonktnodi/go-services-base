package services

import (
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/router"
)

var apiEndpoints *router.Endpoints = nil

func InitServices(filename string) {
  if filename == "" {
    filename = "api-setting"
  }
  apiEndpoints = LoadEndpoints(filename)
}

func GetEndpoint(path, method string) (endpoint *router.EndpointConfig) {
  endpoint = apiEndpoints.GetEndpoint(path, method)
  if endpoint == nil {
    logging.Errorf("failed to find endpoint: [%s]%s", method, path)
    return
  }
  return
}
