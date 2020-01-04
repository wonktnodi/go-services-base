package routers

import (
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/auth"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "time"
)

func authenticator(c *gin.Context) (data interface{}, err error) {
  return
}

func unauthorized(c *gin.Context, code int, message string) {
  return
}

func loginResponse(*gin.Context, int, string, time.Time, interface{}) {
  return
}

func InitSession() auth.AuthorizationHandler {
  authHandler := auth.NewBasicAuthHandler(authenticator, unauthorized, loginResponse)
  if authHandler == nil {
    logging.Fatalf("failed to initialize authorization module")
  }
  return authHandler
}
