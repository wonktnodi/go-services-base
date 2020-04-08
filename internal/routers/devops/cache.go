package devops

import (
  "encoding/binary"
  "github.com/dgryski/go-farm"
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/cache"
  //"github.com/wonktnodi/go-services-base/pkg/encoding/base36"
  "github.com/wonktnodi/go-services-base/pkg/encoding/base62"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/persistence"
  "net/http"
)

func CacheClear() {
  cache.CacheStore.Flush()
}

//func Encode(dst, src []byte) int {
//  j := 0
//  for _, v := range src {
//    dst[j] = hextable[v>>4]
//    dst[j+1] = hextable[v&0x0f]
//    j += 2
//  }
//  return len(src) * 2
//}

func CacheTest(c *gin.Context) {
  cache.CacheStore.Set("test", "test string", persistence.DEFAULT)

  data := make([]byte, binary.MaxVarintLen64)
  len := binary.PutUvarint(data, 12)
  data = append(data[:len], []byte("test string")...)
  id := farm.Hash64(data)
  dst := base62.Encode(id)
  logging.Trace(dst)
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
