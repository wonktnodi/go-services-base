package middleware

import (
  "github.com/gin-gonic/gin"
  "go-services-base/pkg/logging"
  "go-services-base/pkg/restful"
)

func Authorization() gin.HandlerFunc {
  return func(c *gin.Context) { // check authorization here
    sessionToken := restful.GetSessionInfo(c, restful.SESSION_COOKIE_KEY_TOKEN)
    val := sessionToken.Get(restful.SESSION_NAME_TOKEN)
    logging.Trace(val)
    c.Next()
  }
}
