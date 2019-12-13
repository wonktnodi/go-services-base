package middleware

import "github.com/gin-gonic/gin"

func PermissionsMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) { // check authorization here
    c.Next()
  }
}
