package utils

import (
  "fmt"
  "os"
)

func Exit(format string, v ...interface{}) {
  fmt.Printf(format, v...)
  os.Exit(1)
}
