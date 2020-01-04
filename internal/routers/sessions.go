package routers

import (
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/auth"
  "github.com/wonktnodi/go-services-base/pkg/errors"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/restful"
  "time"
)

func authenticator(c *gin.Context) (data interface{}, err error) {
  return
}

func unauthorized(c *gin.Context, code int, message string) {
  return
}

func loginResponse(c *gin.Context, code int, token string, expire time.Time, data interface{}) {
  session := restful.NewApiRequest(c, nil)
  session.Response(code, errors.SUCCESS, data, nil)
  
  return
}

func InitSession() auth.AuthorizationHandler {
  authHandler := auth.NewBasicAuthHandler(authenticator, unauthorized, loginResponse)
  if authHandler == nil {
    logging.Fatalf("failed to initialize authorization module")
  }
  return authHandler
}
