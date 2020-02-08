package utils

func Uint64ArrayRemoveDuplicated(arr []uint64) (ret []uint64) {
  var list = map[uint64]bool{}
  for _, v := range arr {
    list[v] = true
  }
  
  ret = make([]uint64, 0)
  for v, _ := range list {
    ret = append(ret, v)
  }
  return
}
