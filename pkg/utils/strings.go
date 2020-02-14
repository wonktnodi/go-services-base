package utils

import (
  "bytes"
  "github.com/wonktnodi/go-services-base/pkg/errors"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "strconv"
  "strings"
)

func UInt64ArrayToString(A []uint64, delim string) string {
  var buffer bytes.Buffer
  for i := 0; i < len(A); i++ {
    buffer.WriteString(strconv.FormatUint(A[i], 10))
    if i != len(A)-1 {
      buffer.WriteString(delim)
    }
  }
  
  return buffer.String()
}

func StringToUint64Array(val, delim string) (ret []uint64, code int) {
  valArray := strings.Split(val, delim)
  list := make([]uint64, 0)
  for _, v := range valArray {
    d, err := strconv.ParseUint(strings.Trim(v, " "), 10, 64)
    if err != nil {
      logging.Warnf("failed to parse string to uint64, %s", err)
      code = errors.ERROR_PARTICAL_FAIL
    } else {
      list = append(list, d)
    }
  }
  ret = list
  return
}

func ReduceArray(arr []uint64) []uint64 {
  ret := []uint64{}
  
  // remove duplicates
  list := map[uint64]bool{}
  for _, v := range arr {
    list[v] = true
  }
  
  for v, _ := range list {
    ret = append(ret, v)
  }
  
  return ret
}
