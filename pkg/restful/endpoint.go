package restful

import (
  "fmt"
  "net/http"
)

func (e *Endpoint) BuildUrl(vars ...fmt.Stringer) (url string) {
  var rawURL = []byte{}

  rawURL = append(rawURL, e.Backend[0].Host[0]...)
  backendUrl := e.pat.BuildUrl(e.Backend[0].URLPattern, vars...)
  rawURL = append(rawURL, backendUrl...)
  url = string(rawURL)
  return
}

func (e *Endpoint) Get(r *http.Request, cookies map[string]string, vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  url := e.BuildUrl(vars...)
  var queryString = ""
  if e.BringQuery {
    queryString = r.URL.RawQuery
  }
  ret, code = Get(url, queryString, cookies, e.Timeout)
  return
}

func (e *Endpoint) PostJson(r *http.Request, data interface{}, cookies map[string]string,
  vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  url := e.BuildUrl(vars...)
  var queryString = ""
  if e.BringQuery {
    queryString = r.URL.RawQuery
  }
  ret, code = PostJson(url, data, queryString, cookies, e.Timeout)
  return
}

func (e *Endpoint) PutJson(r *http.Request, data interface{}, cookies map[string]string,
  vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  url := e.BuildUrl(vars...)
  var queryString = ""
  if e.BringQuery {
    queryString = r.URL.RawQuery
  }
  ret, code = PutJson(url, data, queryString, cookies, e.Timeout)
  return
}

func (e *Endpoint) Delete(r *http.Request, data interface{}, cookies map[string]string,
  vars ...fmt.Stringer) (ret *BackendResponse, code int) {
  url := e.BuildUrl(vars...)

  var queryString = ""
  if e.BringQuery {
    queryString = r.URL.RawQuery
  }
  ret, code = Delete(url, queryString, data, cookies, e.Timeout)
  return
}
