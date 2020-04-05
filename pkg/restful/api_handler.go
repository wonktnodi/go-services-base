package restful

import (
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/wonktnodi/go-services-base/pkg/errors"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "net/http"
)

type ApiRequest struct {
  gin     *gin.Context
  cookies map[string]string
}

func NewApiRequest(c *gin.Context, cookies map[string]string) *ApiRequest {
  return &ApiRequest{
    gin:     c,
    cookies: cookies,
  }
}

func (r *ApiRequest) GetRawBodyAsJson() (body json.RawMessage, code int) {
  err := r.gin.ShouldBind(&body)
  if err != nil {
    logging.Warnf("parse raw body failed, %s", err)
    code = errors.INTERNAL_ERROR
    return
  }
  return
}

func (r *ApiRequest) Delete(path string, data interface{}, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  endpoint := GetEndpoint(path, "DELETE")
  if endpoint == nil {
    logging.Errorf("can't find router in api endpoint list: [%s]%s", r.gin.Request.Method, path)
    code = errors.INTERNAL_ERROR
    return
  }

  ret, code = endpoint.Delete(r.gin.Request, data, r.cookies, vars...)

  return
}

func (r *ApiRequest) PutJson(path string, data interface{}, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  endpoint := GetEndpoint(path, "PUT")
  if endpoint == nil {
    logging.Errorf("can't find router in api endpoint list: [%s]%s", r.gin.Request.Method, path)
    code = errors.INTERNAL_ERROR
    return
  }
  ret, code = endpoint.PutJson(r.gin.Request, data, r.cookies, vars...)
  return
}

func (r *ApiRequest) PostJson(path string, data interface{}, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  endpoint := GetEndpoint(path, "POST")
  if endpoint == nil {
    logging.Errorf("can't find router in api endpoint list: [%s]%s", r.gin.Request.Method, path)
    code = errors.INTERNAL_ERROR
    return
  }
  ret, code = endpoint.PostJson(r.gin.Request, data, r.cookies, vars...)
  return
}

func (r *ApiRequest) Get(path string, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  endpoint := GetEndpoint(path, "GET")
  if endpoint == nil {
    logging.Errorf("can't find router[%s] in router map", path)
    code = errors.INTERNAL_ERROR
    return
  }
  ret, code = endpoint.Get(r.gin.Request, r.cookies, vars...)
  return
}

func (r *ApiRequest) Response(httpCode, code int, data interface{}, paging *Pagination) {
  var resp Response
  resp.Code = code
  if data != nil {
    resp.Data = data
  }
  if paging != nil {
    resp.Paging = paging
  }
  r.gin.JSON(httpCode, &resp)
}

func (r *ApiRequest) ResponseCode(code int) {
  var resp Response
  resp.Code = code
  r.gin.JSON(http.StatusOK, &resp)
}

func (r *ApiRequest) ResponseData(code int, data interface{}) {
    resp := Response{}
    resp.Code = code
    if data != nil {
        resp.Data = data
    }
    r.gin.JSON(http.StatusOK, &resp)
}

func (r *ApiRequest) SuccessData(data interface{}) {
  resp := Response{}
  if data != nil {
    resp.Data = data
  }
  r.gin.JSON(http.StatusOK, &resp)
}

func (r *ApiRequest) FailedData(code int, data interface{}) {
  resp := Response{}
  resp.Code = code
  if data != nil {
    resp.Data = data
  }
  r.gin.JSON(http.StatusOK, &resp)
}

func (r *ApiRequest) Success(ret *Response) {
  if ret == nil {
    ret = &Response{}
  }
  ret.Code = errors.SUCCESS
  r.gin.JSON(http.StatusOK, ret)
}

func (r *ApiRequest) Failed(ret *Response, code int) {
  if ret == nil {
    ret = &Response{}
  }
  ret.Code = code
  r.gin.JSON(http.StatusOK, ret)
}
