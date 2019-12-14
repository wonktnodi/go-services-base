package config

type ServerSetting struct {
  RunMode      int
  LogLevel     int
  Address      string
  Port         int
  ReadTimeout  int64 // milliseconds
  WriteTimeout int64 // milliseconds
}

type Database struct {
  Type        string
  User        string
  Password    string
  Host        string
  Port        int
  Name        string
  TablePrefix string
}

type RedisSetting struct {
  Address  string
  Port     int
  DB       int
  Pwd      string
  UserName string
}

type LruSetting struct {
}
