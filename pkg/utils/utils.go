package utils

import (
  "fmt"
  "os"
  "reflect"
)

func Exit(format string, v ...interface{}) {
  fmt.Printf(format, v...)
  os.Exit(1)
}

func IsNil(v interface{}) bool {
  return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
