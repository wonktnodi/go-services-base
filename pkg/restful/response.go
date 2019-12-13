package restful

import "github.com/gin-gonic/gin"

type Transaction struct {
  C *gin.Context
}

func (r *Transaction) ResponseCode(httpCode, code int) {
  var resp Response
  resp.Code = code
  r.C.JSON(httpCode, &resp)
}

func (r *Transaction) Response(httpCode, code int, data interface{}) {
  var resp Response
  resp.Code = code
  resp.Data = data
  r.C.JSON(httpCode, &resp)
}

func (r *Transaction) FailedResult(httpCode, code int) {
  var resp Response
  resp.Code = code
  r.C.JSON(httpCode, &resp)
}
