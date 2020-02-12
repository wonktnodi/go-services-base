package demo

import (
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/restful"
)

func GetUsers(c *gin.Context) {
  ret := restful.ParsePagination(c, true)
  logging.Trace(ret)
}
