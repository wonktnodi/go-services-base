package restful

import "encoding/json"

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
  Limit  int `json:"limit"`
  Offset int `json:"offset"`
  Total  int `json:"total"`
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
