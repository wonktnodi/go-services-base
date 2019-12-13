package router

import (
  "fmt"
  "golang.org/x/text/encoding"
  "go-services-base/pkg/logging"
  "time"
)

type ServiceConfig struct {
  // default timeout
  Timeout time.Duration
  // default TTL for GET
  CacheTTL time.Duration
  // default set of hosts
  Host []string
  // version code of the configuration
  Version int
  
  // run in debug mode
  Debug bool
  
  // set of endpoint definitions
  Endpoints []*EndpointConfig
}

type EndpointConfig struct {
  // url pattern to be registered and exposed to the world
  Endpoint string
  // HTTP method of the endpoint (GET, POST, PUT, etc)
  Method string
  // set of definitions of the backend to be linked to this endpoint
  Backend []*Backend
  // number of concurrent calls this endpoint must send to the backend
  ConcurrentCalls int
  // timeout of this endpoint
  Timeout time.Duration
  // duration of the cache header
  CacheTTL time.Duration
  // pass query string to backend
  BringQuery bool
  // list of query string params to be extracted from the URI
  QueryString []string
  // Endpoint Extra configuration for customized behaviour
  ExtraConfig ExtraConfig
  
  pat *Pattern
}

type Backend struct {
  // HTTP method of the request to send to the backend
  Method string
  // Set of hosts of the API
  Host []string
  // URL pattern to use to locate the resource to be consumed
  URLPattern string
  // the encoding format
  Encoding string
  // the response to process is a collection
  IsCollection bool
  // name of the field to extract to the root. If empty, the formatter will do nothing
  Target string
  
  // list of keys to be replaced in the URLPattern
  URLKeys []string
  // number of concurrent calls this endpoint must send to the API
  ConcurrentCalls int
  // timeout of this backend
  Timeout time.Duration
  // decoder to use in order to parse the received response from the API
  Decoder encoding.Decoder
  // Backend Extra configuration for customized behaviours
  ExtraConfig ExtraConfig
}

type ExtraConfig map[string]interface{}

type ParamString string

func (p ParamString) String() string {
  return string(p)
}

type ParamInt64 int64

func (p ParamInt64) String() string {
  return fmt.Sprintf("%d", p)
}

type ParamUint64 uint64

func (p ParamUint64) String() string {
  return fmt.Sprintf("%d", p)
}

func (endpoint *EndpointConfig) BuildUrl(vars ...fmt.Stringer) (url string) {
  var rawURL = []byte{}
  
  rawURL = append(rawURL, endpoint.Backend[0].Host[0]...)
  backendUrl := endpoint.pat.BuildUrl(endpoint.Backend[0].URLPattern, vars...)
  rawURL = append(rawURL, backendUrl...)
  url = string(rawURL)
  return
}

type Endpoints struct {
  config  ServiceConfig
  mapping map[string]*EndpointConfig
}

func NewEndpoints(config *ServiceConfig) *Endpoints {
  mapping := loadEndpoints(config)
  return &Endpoints{
    config:  *config,
    mapping: mapping,
  }
}

func loadEndpoints(config *ServiceConfig) (ret map[string]*EndpointConfig) {
  ret = map[string]*EndpointConfig{}
  
  for _, item := range config.Endpoints {
    // verify endpoint
    if len(item.Backend) == 0 {
      continue
    }
    if item.Timeout == 0 {
      item.Timeout = config.Timeout
    }
    item.pat = NewUriPattern(item.Endpoint)
    ret[fmt.Sprintf("%s_%s", item.Method, item.Endpoint)] = item
  }
  return
}

func (e *Endpoints) GetEndpoint(endpoint, method string) (ret *EndpointConfig) {
  if e.mapping == nil {
    return
  }
  
  ret, ok := e.mapping[fmt.Sprintf("%s_%s", method, endpoint)]
  if ok == false {
    logging.Errorf("failed to get endpoint [%s]%s", method, endpoint)
    ret = nil
  }
  return
}
