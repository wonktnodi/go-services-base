package auth

import "context"

type SessionHandShake struct {
  Alg    string `json:"alg,omitempty"`
  Salt   string `json:"salt,omitempty"`
  Expire int    `json:"expire,omitempty"` // milliseconds
}

type SignInfo struct {
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
