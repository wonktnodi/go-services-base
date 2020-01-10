package routers

import (
  "encoding/gob"
  "errors"
  "github.com/gin-gonic/gin"
  "github.com/mitchellh/mapstructure"
  "github.com/wonktnodi/go-services-base/pkg/auth"
  serviceErrors "github.com/wonktnodi/go-services-base/pkg/errors"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/restful"
  "github.com/wonktnodi/go-services-base/pkg/sessions"
  "time"
)

type userInfo struct {
  UserName string `json:"userName"`
  UserId   uint64 `json:"userId"`
}

func authenticator(c *gin.Context, info auth.SignInInfo) (data interface{}, code int) {
  store := sessions.DefaultMany(c, restful.SESSION_COOKIE_KEY_CODE)
  detail := store.Get("content")
  if detail == nil {
    code = serviceErrors.INVALID_PARAMS
    return
  }
  var savedInfo VerificationInfo
  err := mapstructure.Decode(detail, &savedInfo)
  if err != nil {
    logging.Errorf("failed to parse interface to data, %s", err)
    code = serviceErrors.INTERNAL_ERROR
    return
  }
  if savedInfo.Code != info.Code {
    err = errors.New("login failed")
    code = serviceErrors.ERROR_AUTH
    return
  }
  
  var options = sessions.Options{
    MaxAge:   -1,
    HttpOnly: true,
    Path:     "/",
  }
  store.Clear()
  store.Options(options)
  store.Save()
  
  // AllowAllOrigins
  data = "test data"
  return
}

func unauthorized(c *gin.Context, code int, message string) {
  session := restful.NewApiRequest(c, nil)
  session.ResponseCode(code)
  return
}

func loginResponse(c *gin.Context, code int, token string, expire time.Time, data interface{}) {
  sessionStore := sessions.DefaultMany(c, restful.SESSION_COOKIE_KEY_TOKEN)
  sessionStore.Set("token", data)
  sessionStore.Save()
  
  session := restful.NewApiRequest(c, nil)
  session.Response(code, serviceErrors.SUCCESS, data, nil)
  
  return
}

type VerificationInfo struct {
  Type      int
  Code      string
  Timestamp int64
}

func generateVerifyCode(c *gin.Context, info auth.SignInInfo) (ret interface{}, code int) {
  var data = VerificationInfo{
    Type: info.Type, Code: "111111", Timestamp: time.Now().Unix(),
  }
  store := sessions.DefaultMany(c, restful.SESSION_COOKIE_KEY_CODE)
  var options = store.GetOptions()
  options.MaxAge = 5 * 60
  store.Options(options)
  
  store.Set("content", data)
  err := store.Save()
  if err != nil {
    logging.Warnf("failed to save session info[%s], %s", restful.SESSION_COOKIE_KEY_CODE, err)
    code = serviceErrors.INTERNAL_ERROR
  }
  
  ret = data
  return
}

func InitSession() auth.AuthorizationHandler {
  gob.Register(VerificationInfo{})
  
  authHandler := auth.NewBasicAuthHandler(authenticator, unauthorized, generateVerifyCode, loginResponse)
  if authHandler == nil {
    logging.Fatalf("failed to initialize authorization module")
  }
  return authHandler
}
