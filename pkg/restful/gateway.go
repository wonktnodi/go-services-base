package restful

import (
  "github.com/wonktnodi/go-services-base/pkg/logging"
)

var apiEndpoints *Endpoints = nil

func InitServices(filename string) {
  if filename == "" {
    filename = "api-setting"
  }
  apiEndpoints = LoadEndpoints(filename)
}

func GetEndpoint(path, method string) (endpoint *Endpoint) {
  endpoint = apiEndpoints.GetEndpoint(path, method)
  if endpoint == nil {
    logging.Errorf("failed to find endpoint: [%s]%s", method, path)
    return
  }
  return
}
