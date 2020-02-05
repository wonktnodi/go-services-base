package restful

import (
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/go-resty/resty/v2"
  "github.com/wonktnodi/go-services-base/pkg/errors"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/services"
  "net/http"
  "net/http/httputil"
)

type ApiRequest struct {
  gin *gin.Context
  //endpoint *router.EndpointConfig
  cookies map[string]string
}

func NewApiRequest(c *gin.Context, cookies map[string]string) *ApiRequest {
  return &ApiRequest{
    gin: c,
    //endpoint: endpoint,
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

func (r *ApiRequest) Delete(path string, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  endpoint := services.GetEndpoint(path, "DELETE")
  //endpoint := services.GetEndpoint(path, r.gin.Request.Method)
  if endpoint == nil {
    logging.Errorf("can't find router in api endpoint list: [%s]%s", r.gin.Request.Method, path)
    code = errors.INTERNAL_ERROR
    return
  }
  url := endpoint.BuildUrl(vars...)
  client := resty.New()
  if endpoint.Timeout != 0 {
    client.SetTimeout(endpoint.Timeout)
  }
  
  var response BackendResponse
  request := client.R().EnableTrace()
  if endpoint.BringQuery {
    request.SetQueryString(r.gin.Request.URL.RawQuery)
  }
  for k, c := range r.cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }
  
  resp, err := request.Delete(url)
  if err != nil {
    logging.Errorf("%s: failed to send post data, %s", url, err)
    code = errors.INTERNAL_ERROR
    return
  }
  dumpRequest, _ := httputil.DumpRequest(request.RawRequest, true)
  requestBody, _ := json.Marshal(request.Body)
  logging.Tracef("Sending ===>:\n %s\n%+v\nReceived ===>:\n %s",
    string(dumpRequest), string(requestBody), resp.String())
  
  err = json.Unmarshal(resp.Body(), &response)
  if err != nil {
    logging.Warnf("failed to parse [%s]%s response body: %s", r.gin.Request.Method, path, err)
  }
  //if response.Code != errors.SUCCESS {
  //  code = errors.INTERNAL_ERROR
  //}
  ret = &response
  return
}

func (r *ApiRequest) Put(path string, data interface{}, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  //endpoint := services.GetEndpoint(path, r.gin.Request.Method)
  endpoint := services.GetEndpoint(path, "PUT")
  if endpoint == nil {
    logging.Errorf("can't find router in api endpoint list: [%s]%s", r.gin.Request.Method, path)
    code = errors.INTERNAL_ERROR
    return
  }
  url := endpoint.BuildUrl(vars...)
  client := resty.New()
  if endpoint.Timeout != 0 {
    client.SetTimeout(endpoint.Timeout)
  }
  
  var response BackendResponse
  request := client.R().EnableTrace()
  if endpoint.BringQuery {
    request.SetQueryString(r.gin.Request.URL.RawQuery)
  }
  for k, c := range r.cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }
  
  resp, err := request.SetBody(data).Put(url)
  if err != nil {
    logging.Errorf("%s: failed to send post data, %s", url, err)
    code = errors.INTERNAL_ERROR
    return
  }
  dumpRequest, _ := httputil.DumpRequest(request.RawRequest, true)
  requestBody, _ := json.Marshal(request.Body)
  logging.Tracef("Sending ===>:\n %s\n%+v\nReceived ===>:\n %s",
    string(dumpRequest), string(requestBody), resp.String())
  
  err = json.Unmarshal(resp.Body(), &response)
  if err != nil {
    logging.Warnf("failed to parse [%s]%s response body: %s", r.gin.Request.Method, path, err)
  }
  //if response.Code != errors.SUCCESS {
  //  code = errors.INTERNAL_ERROR
  //}
  ret = &response
  return
}

func (r *ApiRequest) Post(path string, data interface{}, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  //endpoint := services.GetEndpoint(path, r.gin.Request.Method)
  endpoint := services.GetEndpoint(path, "POST")
  if endpoint == nil {
    logging.Errorf("can't find router in api endpoint list: [%s]%s", r.gin.Request.Method, path)
    code = errors.INTERNAL_ERROR
    return
  }
  url := endpoint.BuildUrl(vars...)
  client := resty.New()
  if endpoint.Timeout != 0 {
    client.SetTimeout(endpoint.Timeout)
  }
  var response BackendResponse
  request := client.R().EnableTrace()
  if endpoint.BringQuery {
    request.SetQueryString(r.gin.Request.URL.RawQuery)
  }
  for k, c := range r.cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }
  
  resp, err := request.SetBody(data).Post(url)
  if err != nil {
    logging.Errorf("%s: failed to send post data, %s", url, err)
    code = errors.INTERNAL_ERROR
    return
  }
  dumpRequest, _ := httputil.DumpRequest(request.RawRequest, true)
  requestBody, _ := json.Marshal(request.Body)
  logging.Tracef("Sending ===>:\n %s\n%+v\nReceived ===>:\n %s",
    string(dumpRequest), string(requestBody), resp.String())
  
  err = json.Unmarshal(resp.Body(), &response)
  if err != nil {
    logging.Warnf("failed to parse [%s]%s response body: %s", r.gin.Request.Method, path, err)
  }
  //if response.Code != errors.SUCCESS {
  //  code = errors.INTERNAL_ERROR
  //}
  ret = &response
  return
}

func (r *ApiRequest) Get(path string, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  //endpoint := services.GetEndpoint(path, r.gin.Request.Method)
  endpoint := services.GetEndpoint(path, "GET")
  if endpoint == nil {
    logging.Errorf("can't find router[%s] in router map", path)
    code = errors.INTERNAL_ERROR
    return
  }
  
  url := endpoint.BuildUrl(vars...)
  client := resty.New()
  if endpoint.Timeout != 0 {
    client.SetTimeout(endpoint.Timeout)
  }
  var response BackendResponse
  request := client.R().EnableTrace()
  if endpoint.BringQuery {
    request.SetQueryString(r.gin.Request.URL.RawQuery)
  }
  for k, c := range r.cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }
  
  resp, err := request.SetResult(&response).Get(url)
  if resp.StatusCode() != http.StatusOK {
    code = resp.StatusCode()
    return
  }
  dumpRequest, _ := httputil.DumpRequest(request.RawRequest, true)
  logging.Tracef("Sending ===>:\n %s\nReceived ===>:\n %s",
    string(dumpRequest), resp.String())
  
  if err != nil {
    logging.Errorf("%s: failed to get response, %s", url, err)
    code = errors.INTERNAL_ERROR
    return
  }
  ret = &response
  if ret.Code != errors.SUCCESS {
    logging.Warnf("[GET]%s failed with code %d", url, ret.Code)
    //code = errors.INTERNAL_REQUEST_FAILED
  }
  return
}

func (r *ApiRequest) Response(httpCode, code int, data interface{}, paging *Pagination) {
  var resp Response
  resp.Code = code
  resp.Data = data
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

func (r *ApiRequest) FailedResult(code int) {
  var resp Response
  resp.Code = code
  r.gin.JSON(http.StatusOK, &resp)
}
