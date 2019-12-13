package utils

import (
  "bytes"
  "unsafe"
)

func BytesReplace(s, old, new []byte, n int, exact bool) []byte {
  if n == 0 {
    return s
  }
  
  lenOld := len(old)
  lenNew := len(new)
  //lenStr := len(s)
  if lenOld < lenNew {
    return bytes.Replace(s, old, new, n)
  }
  if exact {
    lenOld = lenOld - lenNew
  }
  if n < 0 {
    n = len(s)
  }
  
  var wid, i, j, w int
  var l = len(s)
  for i, j = 0, 0; i < l && j < n; j++ {
    wid = bytes.Index(s[i:], old)
    if wid < 0 {
      break
    }
    
    w += copy(s[w:], s[i:i+wid])
    w += copy(s[w:], new)
    if exact {
      if l > w+lenOld {
        copy(s[w:], s[w+lenOld:])
      }
      s = s[:l-lenOld]
      l -= lenOld
    }
    i += wid + len(new)
  }
  
  w += copy(s[w:], s[i:])
  return s[0:w]
}

func BytesToString(s []byte) string {
  return *(*string)(unsafe.Pointer(&s))
}
