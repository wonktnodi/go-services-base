package base62

import "math"

var base = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func Encode(num uint64) string {
  baseStr := ""
  for {
    if num <= 0 {
      break
    }

    i := num % 62
    baseStr += base[i]
    num = (num - i) / 62
  }
  return baseStr
}

func Decode(base62 string) uint64 {
  rs := uint64(0)
  len := len(base62)
  f := flip(base)
  for i := 0; i < len; i++ {
    rs += f[string(base62[i])] * uint64(math.Pow(62, float64(i)))
  }
  return rs
}

func flip(s []string) map[string]uint64 {
  f := make(map[string]uint64)
  for index, value := range s {
    f[value] = uint64(index)
  }
  return f
}
