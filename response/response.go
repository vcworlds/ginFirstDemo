package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
func Success(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Error(c *gin.Context, msg string) {
	Response(c, http.StatusUnprocessableEntity, 422, nil, msg)
}
