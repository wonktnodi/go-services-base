package middleware

import (
  "github.com/gin-gonic/gin"
  "services-base/pkg/logging"
  "services-base/pkg/restful"
)

func Authorization() gin.HandlerFunc {
  return func(c *gin.Context) { // check authorization here
    sessionToken := restful.GetSessionInfo(c, restful.SESSION_COOKIE_KEY_TOKEN)
    val := sessionToken.Get(restful.SESSION_NAME_TOKEN)
    logging.Trace(val)
    c.Next()
  }
}
