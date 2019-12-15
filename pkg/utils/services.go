package utils

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "time"
)

func StartService(router *gin.Engine, setting *config.ServerSetting) error {
  if setting.RunMode == 0 {
    gin.SetMode("debug")
  }
  endPoint := fmt.Sprintf("%s:%d", setting.Address, setting.Port)
  readTimeout := time.Millisecond * time.Duration(setting.ReadTimeout)
  writeTimeout := time.Duration(setting.WriteTimeout) * time.Millisecond
  maxHeaderBytes := 1 << 20
  
  server := &http.Server{
    Addr:           endPoint,
    Handler:        router,
    MaxHeaderBytes: maxHeaderBytes,}
  
  if setting.ReadTimeout > 0 {
    server.ReadTimeout = readTimeout
  }
  if setting.WriteTimeout > 0 {
    server.WriteTimeout = writeTimeout
  }
  
  logging.Infof("start service at %s", endPoint)
  return server.ListenAndServe()
}
