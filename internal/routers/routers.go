package routers

import (
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/internal/routers/devops"
  "github.com/wonktnodi/go-services-base/pkg/restful"
  "time"
)

func InitRouters() *gin.Engine {
  r := gin.New()
  r.Use(gin.Logger()) // 日志
  
  r.Use(cors.New(cors.Config{
    //AllowAllOrigins:  true,
    AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
    AllowHeaders:     []string{"Origin", "content-type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return true
    },
    MaxAge: 12 * time.Hour,
  }))                 // 跨域请求
  r.Use(gin.Recovery())
  
  r.GET("/version", restful.Version)
  r.GET("devops/cache/demo", devops.CacheTest)
  r.GET("devops/cache/demoGet", devops.CacheTestGet)
  
  return r
}
