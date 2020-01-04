package routers

import (
  "encoding/gob"
  "fmt"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/internal/config"
  "github.com/wonktnodi/go-services-base/internal/models"
  "github.com/wonktnodi/go-services-base/internal/routers/devops"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/restful"
  "github.com/wonktnodi/go-services-base/pkg/sessions"
  "github.com/wonktnodi/go-services-base/pkg/sessions/redis"
  "time"
)

var sessionNames = []string{
  restful.SESSION_COOKIE_KEY_TOKEN,
  restful.SESSION_COOKIE_KEY_LOGIN,
  restful.SESSION_COOKIE_KEY_CODE,
  restful.SESSION_COOKIE_KEY_SESSION,
}

func InitRouters() *gin.Engine {
  gob.Register(models.SessionInfo{})
  store, err := redis.NewStore(10, "tcp",
    fmt.Sprintf("%s:%d", config.Settings.Redis.Address, config.Settings.Redis.Port),
    "", []byte("secret"))
  if err != nil {
    logging.Fatalf("failed to create session store: %v", err)
    return nil
  }
  
  authHandler := InitSession()
  r := gin.New()
  r.Use(gin.Logger()) // 日志
  r.Use(sessions.SessionsMany(sessionNames, store))
  
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
  })) // 跨域请求
  r.Use(gin.Recovery())
  
  r.GET("/version", restful.Version)
  r.GET("devops/cache/demo", devops.CacheTest)
  r.GET("devops/cache/demoGet", devops.CacheTestGet)
  
  apiV1 := r.Group("v1")
  apiV1.GET("/sessions/new", authHandler.Handshake)
  apiV1.POST("/sessions/verifyCode", authHandler.GenerateVerifyCode)
  apiV1.POST("/sessions", authHandler.SignIn)
  apiV1.PUT("/sessions", authHandler.RefreshSession)
  apiV1.DELETE("/sessions", authHandler.SignOut)
  
  return r
}
