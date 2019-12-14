package restful

import (
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "strconv"
)

const (
  SESSION_COOKIE_KEY_SESSION = "sid"
  SESSION_COOKIE_KEY_LOGIN   = "lsid"
  SESSION_COOKIE_KEY_CODE    = "msid"
  SESSION_COOKIE_KEY_TOKEN   = "token"
  
  SESSION_NAME_SESSION = "session-info"
  SESSION_NAME_LOGIN   = "login-session"
  SESSION_NAME_TOKEN   = "token"
)

var EmptyData interface{}

type Pagination struct {
  Limit    int `json:"limit"`
  Offset   int `json:"offset"`
  Total    int `json:"total"`
  StartRow int `json:"-"`
}

type BackendResponse struct {
  Code   int             `json:"code"`
  Msg    string          `json:"msg,omitempty"`
  Data   json.RawMessage `json:"data,omitempty"`
  Paging *Pagination     `json:"paging,omitempty"`
}

type Response struct {
  Code   int         `json:"code"`
  Msg    string      `json:"msg,omitempty"`
  Data   interface{} `json:"data,omitempty"`
  Paging *Pagination `json:"paging,omitempty"`
}

func ParsePagination(c *gin.Context) (ret *Pagination) {
  var paging Pagination
  var err error
  strVal := c.Query("limit")
  if strVal == "" {
    paging.Limit = 20
  } else {
    paging.Limit, err = strconv.Atoi(strVal)
    if err != nil {
      logging.WarnF("failed to parse pagination limit, %s", err)
      return nil
    }
  }
  
  strVal = c.Query("offset")
  if strVal == "" {
    paging.Offset = 1
  } else {
    paging.Offset, err = strconv.Atoi(strVal)
    if err != nil {
      logging.WarnF("failed to parse pagination offset, %s", err)
    }
  }
  
  if paging.Offset < 1 {
    paging.Offset = 1
  }
  paging.StartRow = (paging.Offset - 1) * paging.Limit
  
  ret = &paging
  return
}
