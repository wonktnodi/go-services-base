package utils

import (
  "context"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
)

var quit = make(chan os.Signal, 1)

func CheckServiceStatus() {
  _, ok := <-quit

  if ok {
    close(quit)
  }
}

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
  go func() {
    // service connections
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      logging.Fatalf("listen: %s\n", err)
    }
  }()

  // Wait for interrupt signal to gracefully shutdown the server with
  // a timeout of 5 seconds.

  // kill (no param) default send syscall.SIGTERM
  // kill -2 is syscall.SIGINT
  // kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

  CheckServiceStatus()

  log.Println("Shutdown Server ...")

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  if err := server.Shutdown(ctx); err != nil {
    log.Fatal("Server Shutdown: ", err)
  }
  logging.Info("Server exiting")
  return nil
}
