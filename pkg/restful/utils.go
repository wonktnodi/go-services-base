package restful

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "go-services-base/pkg"
  "go-services-base/pkg/sessions"
)

func Version(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "buildTime": pkg.BuildDate,
    "version":   pkg.Version,
    "commit":    pkg.CommitHash})
}

// GetToken 获取用户令牌
func GetSessionInfo(c *gin.Context, name string) sessions.Session {
  return sessions.DefaultMany(c, name)
}
