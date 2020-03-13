package demo

import (
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/restful"
  "net/http"
  "time"
)

func GetUsers(c *gin.Context) {
  //ret := restful.ParsePagination(c, true)
  endpoint := restful.GetEndpoint("/posts", "GET")
  if nil == endpoint {
    return
  }

  ret, _ := endpoint.Get(c.Request, nil, time.Second*5)
  c.JSON(http.StatusOK, ret)
}

func CreateUsers(c *gin.Context) {
  endpoint := restful.GetEndpoint("/posts/:id", "POST")
  if nil == endpoint {
    return
  }

  data := json.RawMessage([]byte(`{
	"uid": 100000025,
	"nick": "sssssss",
	"img": "imag urls",
	"refCode": "adfadfasdfasdfasdfasdf",
	"couponId": 1,
	"refCouponId": 12,
	"operation": 2
}`))
  id := "11111111"
  ret, _ := endpoint.PostJson(c.Request, data, nil, restful.ParamString(id))
  c.JSON(http.StatusOK, ret)
}

func DeleteUsers(c *gin.Context) {
  endpoint := restful.GetEndpoint("/posts/:id", "DELETE")
  if nil == endpoint {
    return
  }

  id := "70"
  ret, _ := endpoint.Delete(c.Request, nil, nil, restful.ParamString(id))
  c.JSON(http.StatusOK, ret)
}
