package models

type SessionInfo struct {
  Alg    string `json:"alg,omitempty"`
  Salt   string `json:"salt,omitempty"`
  Expire int    `json:"expire,omitempty"` // milliseconds
}

type SignInfo struct {
  Id    string `json:"id,omitempty"`
  DevId string `json:"devId,omitempty"`
  Type  int    `json:"type,omitempty"`
  Code  string `json:"code,omitempty"`
}

type VerifyInfo struct {
  Id         uint64 `json:"id"`
  VerifyCode string `json:"verifyCode,omitempty"`
  Uid        uint64 `json:"uid"`
  Expire     int    `json:"expire,omitempty"`
}
