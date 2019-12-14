package devops

import (
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/cache"
  "github.com/wonktnodi/go-services-base/pkg/persistence"
  "net/http"
)

func CacheClear() {
  cache.CacheStore.Flush()
}

func CacheTest(c *gin.Context) {
  cache.CacheStore.Set("test", "test string", persistence.DEFAULT)
  c.Status(http.StatusOK)
}

func CacheTestGet(c *gin.Context) {
  var val string
  err := cache.CacheStore.Get("test", &val)
  if err != nil {
    c.String(http.StatusInternalServerError, err.Error())
  }
  c.String(http.StatusOK, val)
}
