package auth

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/wonktnodi/go-services-base/pkg/errors"
    "github.com/wonktnodi/go-services-base/pkg/logging"
    "github.com/wonktnodi/go-services-base/pkg/restful"
    "github.com/wonktnodi/go-services-base/pkg/sessions"
    "time"

    "net/http"
)

//github.com/gin-contrib/sessions v0.0.3 // indirect

const (
    SESSION_COOKIE_KEY_SESSION = "sid"
    SESSION_COOKIE_KEY_LOGIN   = "lsid"
    SESSION_COOKIE_KEY_CODE    = "msid"
    SESSION_COOKIE_KEY_TOKEN   = "token"
)

type SessionHandShake struct {
    Alg    string `json:"alg,omitempty"`
    Salt   string `json:"salt,omitempty"`
    Expire int    `json:"expire,omitempty"` // milliseconds
}

type SignInInfo struct {
    ID    string `json:"id,omitempty"`
    DevId string `json:"devId,omitempty"`
    Type  int    `json:"type,omitempty"`
    Code  string `json:"code,omitempty"`
}

type VerifyInfo struct {
    ID         uint64 `json:"id"`
    VerifyCode string `json:"verifyCode,omitempty"`
    Uid        uint64 `json:"uid,omitempty"`
    Expire     int    `json:"expire,omitempty"`
}

// TokenInfo 令牌信息
type TokenInfo interface {
    // 获取访问令牌
    GetAccessToken() string
    // 获取令牌类型
    GetTokenType() string
    // 获取令牌到期时间戳
    GetExpiresAt() int64
    // JSON编码
    EncodeToJSON() ([]byte, error)
}

// Author 认证接口
type Author interface {
    // 生成令牌
    GenerateToken(ctx context.Context, userID string) (TokenInfo, error)

    // 销毁令牌
    DestroyToken(ctx context.Context, accessToken string) error

    // 解析用户ID
    ParseUserID(ctx context.Context, accessToken string) (string, error)

    // 释放资源
    Release() error
}

type AuthorizationHandler interface {
    Handshake(c *gin.Context)
    GenerateVerifyCode(c *gin.Context)
    SignIn(c *gin.Context)
    SignOut(c *gin.Context)
    RefreshSession(c *gin.Context)
}

type BasicAuthHandler struct {
    Authenticator func(c *gin.Context, info SignInInfo) (interface{}, int)
    //Authorizer func(data interface{}, c *gin.Context) bool
    Unauthorized        func(c *gin.Context, code int, message string)
    VerifyCodeGenerator func(c *gin.Context, info SignInInfo) (interface{}, int)
    LoginResponse       func(*gin.Context, int, string, time.Time, interface {
    })
}

func NewBasicAuthHandler(
  Authenticator func(c *gin.Context, info SignInInfo) (interface{}, int),
  Unauthorized func(c *gin.Context, code int, message string),
  VerifyCodeGenerator func(c *gin.Context, info SignInInfo) (interface{}, int),
  LoginResponse func(*gin.Context, int, string, time.Time, interface{})) (ret AuthorizationHandler) {
    ret = &BasicAuthHandler{
        Authenticator:       Authenticator,
        Unauthorized:        Unauthorized,
        VerifyCodeGenerator: VerifyCodeGenerator,
        LoginResponse:       LoginResponse,
    }
    return
}

func (h *BasicAuthHandler) Handshake(c *gin.Context) {
    session := restful.NewApiRequest(c, nil)

    sessions.DefaultMany(c, SESSION_COOKIE_KEY_LOGIN)
    sessions.DefaultMany(c, SESSION_COOKIE_KEY_SESSION)
    var info = SessionHandShake{}
    info.Alg = "md5"
    session.Response(http.StatusOK, errors.SUCCESS, info, nil)
}

func (h *BasicAuthHandler) GenerateVerifyCode(c *gin.Context) {
    var info SignInInfo
    err := c.ShouldBind(&info)
    if err != nil {
        logging.Warnf("failed to parse verify info: %s", err)
        c.Status(http.StatusBadRequest)
        return
    }
    session := restful.NewApiRequest(c, nil)
    ret, code := h.VerifyCodeGenerator(c, info)
    if code != errors.SUCCESS {
        session.ResponseCode(code)
        return
    }

    session.ResponseData(code, ret)
}

func (h *BasicAuthHandler) SignIn(c *gin.Context) {
    var info SignInInfo
    err := c.ShouldBind(&info)
    if err != nil {
        logging.Warnf("failed to parse verify info: %s", err)
        c.Status(http.StatusBadRequest)
        return
    }

    data, code := h.Authenticator(c, info)
    if code != errors.SUCCESS {
        h.Unauthorized(c, code, "")
        return
    }

    h.LoginResponse(c, http.StatusOK, "", time.Now(), data)
}

func (h *BasicAuthHandler) SignOut(c *gin.Context) {
    sessionToken := sessions.DefaultMany(c, SESSION_COOKIE_KEY_TOKEN)
    sessionSid := sessions.DefaultMany(c, SESSION_COOKIE_KEY_SESSION)

    var options = sessions.Options{
        MaxAge:   -1,
        HttpOnly: true,
        Path:     "/",
    }

    sessionToken.Clear()
    sessionToken.Options(options)
    sessionToken.Save()
    sessionSid.Clear()
    sessionSid.Options(options)
    sessionSid.Save()

    c.Status(http.StatusNoContent)
}

func (h *BasicAuthHandler) RefreshSession(c *gin.Context) {
    sessionStore := sessions.DefaultMany(c, restful.SESSION_COOKIE_KEY_TOKEN)
    data := sessionStore.Get(restful.SESSION_COOKIE_KEY_TOKEN)
    sessionStore.Set(restful.SESSION_COOKIE_KEY_TOKEN, data)
    sessionStore.Save()

    session := restful.NewApiRequest(c, nil)
    session.ResponseCode(errors.SUCCESS)
    //restful.ResponseData()
    //c.Status(http.StatusOK)
}
