package services

import (
  "services-base/pkg/logging"
  "services-base/pkg/router"
)

var apiEndpoints *router.Endpoints = nil

func InitServices() {
  apiEndpoints = LoadEndpoints("api-setting")
}

func GetEndpoint(path, method string) (endpoint *router.EndpointConfig) {
  endpoint = apiEndpoints.GetEndpoint(path, method)
  if endpoint == nil {
    logging.Errorf("failed to find endpoint: [%s]%s", method, path)
    return
  }
  return
}
