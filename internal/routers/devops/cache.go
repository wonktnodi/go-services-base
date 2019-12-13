package devops

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "go-services-base/internal/models"
  "go-services-base/pkg/persistence"
)

func CacheClear() {
  models.CacheStore.Flush()
}

func CacheTest(c *gin.Context) {
  models.CacheStore.Set("test", "test string", persistence.DEFAULT)
  c.Status(http.StatusOK)
}

func CacheTestGet(c *gin.Context) {
  var val string
  err := models.CacheStore.Get("test", &val)
  if err != nil {
    c.String(http.StatusInternalServerError, err.Error())
  }
  c.String(http.StatusOK, val)
}
