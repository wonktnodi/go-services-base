package restful

import (
  "encoding/json"
  "github.com/go-resty/resty/v2"
  "github.com/wonktnodi/go-services-base/pkg/errors"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "net/http"
  "net/http/httputil"
  "time"
)

var defaultTimeout = 0
var debug = true

func PutForm(url string, data interface{}, rawQuery string,
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  return put(url, data, true, rawQuery, cookies, timeout)
}

func PutJson(url string, data interface{}, rawQuery string,
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  return put(url, data, false, rawQuery, cookies, timeout)
}

func put(url string, data interface{}, form bool, rawQuery string,
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  client := resty.New()
  client.SetDebug(debug)
  client.SetLogger(logging.GetLogger())
  if timeout != 0 {
    client.SetTimeout(timeout)
  }
  var response BackendResponse
  request := client.R().EnableTrace()
  if rawQuery != "" {
    request.SetQueryString(rawQuery)
  }
  for k, c := range cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }

  if form == true {
    request = request.SetFormData(data.(map[string]string))
  } else {
    request = request.SetBody(data)
  }

  resp, err := request.Put(url)
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
    logging.Warnf("failed to parse [PUT]%s response body: %s", url, err)
  }
  ret = &response
  return
}

func PostForm(url string, data interface{}, rawQuery string,
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  return post(url, data, true, rawQuery, cookies, timeout)
}

func PostJson(url string, data interface{}, rawQuery string,
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  return post(url, data, false, rawQuery, cookies, timeout)
}

func post(url string, data interface{}, form bool, rawQuery string,
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  client := resty.New()
  client.SetDebug(debug)
  client.SetLogger(logging.GetLogger())

  if timeout != 0 {
    client.SetTimeout(timeout)
  }
  var response BackendResponse
  request := client.R().EnableTrace()
  if rawQuery != "" {
    request.SetQueryString(rawQuery)
  }
  for k, c := range cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }
  if form == true {
    request = request.SetFormData(data.(map[string]string))
  } else {
    request = request.SetBody(data)
  }

  resp, err := request.Post(url)
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
    logging.Warnf("failed to parse [POST]%s response body: %s", url, err)
  }
  ret = &response
  return
}

func Get(url string, rawQuery string, data interface{},
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  client := resty.New()
  client.SetDebug(debug)
  client.SetLogger(logging.GetLogger())

  if timeout != 0 {
    client.SetTimeout(timeout)
  }

  var response BackendResponse
  request := client.R().EnableTrace()
  if rawQuery != "" {
    request.SetQueryString(rawQuery)
  }
  for k, c := range cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }

  resp, err := request.Get(url)
  if resp.StatusCode() != http.StatusOK {
    code = resp.StatusCode()
    return
  }
  if err != nil {
    logging.Errorf("%s: failed to get response, %s", url, err)
    code = errors.INTERNAL_ERROR
    return
  }
  dumpRequest, _ := httputil.DumpRequest(request.RawRequest, true)
  requestBody, _ := json.Marshal(request.Body)
  logging.Tracef("Sending ===>:\n %s\n%+v\nReceived ===>:\n %s",
    string(dumpRequest), string(requestBody), resp.String())
  err = json.Unmarshal(resp.Body(), &response)
  if err != nil {
    logging.Warnf("failed to parse [GET]%s response body: %s", url, err)
  }
  ret = &response
  return
}

func Delete(url string, rawQuery string, data interface{},
  cookies map[string]string, timeout time.Duration) (ret *BackendResponse, code int) {
  client := resty.New()
  client.SetDebug(debug)
  client.SetLogger(logging.GetLogger())

  if timeout != 0 {
    client.SetTimeout(timeout)
  }

  var response BackendResponse
  request := client.R().EnableTrace()
  if rawQuery != "" {
    request.SetQueryString(rawQuery)
  }
  for k, c := range cookies {
    request.SetCookie(&http.Cookie{
      Name:  k,
      Value: c,
    })
  }

  if data != nil {
    request.SetBody(data)
  }
  resp, err := request.Delete(url)
  if err != nil {
    logging.Errorf("%s: failed to delete[%s], %s", url, err)
    code = errors.INTERNAL_ERROR
    return
  }
  dumpRequest, _ := httputil.DumpRequest(request.RawRequest, true)
  requestBody, _ := json.Marshal(request.Body)
  logging.Tracef("Sending ===>:\n %s\n%+v\nReceived ===>:\n %s",
    string(dumpRequest), string(requestBody), resp.String())
  err = json.Unmarshal(resp.Body(), &response)
  if err != nil {
    logging.Warnf("failed to parse [DELETE]%s response body: %s", url, err)
  }
  ret = &response
  return
}
