package middleware

import (
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/restful"
  "net/http"
)

type SessionUnauthorized func(c *gin.Context)
type SessionAuthorized func(c *gin.Context, data interface{})

func Authorization(authorized SessionAuthorized, unAuthorized SessionUnauthorized) gin.HandlerFunc {
  return func(c *gin.Context) { // check authorization here
    sessionToken := restful.GetSessionInfo(c, restful.SESSION_COOKIE_KEY_TOKEN)
    if sessionToken == nil {
      if unAuthorized != nil {
        unAuthorized(c)
      } else {
        c.Status(http.StatusUnauthorized)
      }
      c.Abort()
      return
    }
    val := sessionToken.Get(restful.SESSION_NAME_TOKEN)
    if nil == val {
      if unAuthorized != nil {
        unAuthorized(c)
      } else {
        c.Status(http.StatusUnauthorized)
      }
      c.Abort()
      return
    }
    if authorized != nil {
      authorized(c, val)
    }
    logging.Trace(val)

    c.Next()
  }
}
