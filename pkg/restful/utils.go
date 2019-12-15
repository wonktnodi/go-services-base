package restful

import (
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg"
  "github.com/wonktnodi/go-services-base/pkg/errors"
  "github.com/wonktnodi/go-services-base/pkg/sessions"
  "net/http"
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

func ResponseData(session *ApiRequest, code int, data interface{}) {
  if code != errors.SUCCESS {
    session.FailedResult(code)
    return
  }
  resp := Response{
    Data: data,
  }
  session.Success(&resp)
}

func ResponseDataWithPagination(session *ApiRequest, code int, data interface{}, paging *Pagination) {
  if code != errors.SUCCESS {
    session.FailedResult(code)
    return
  }
  
  resp := Response{
    Data: data,
  }
  if paging != nil {
    resp.Paging = paging
  }
  session.Success(&resp)
}
