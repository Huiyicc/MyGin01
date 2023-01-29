package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnifiedResponse(ctx *gin.Context, httpStatus int, code int, msg string)  {
	ctx.JSON(httpStatus,gin.H{"code":code, "msg":msg})
}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string)   {
	ctx.JSON(httpStatus,gin.H{"code":code, "data":data,"msg":msg})
}
func Success(ctx *gin.Context,data gin.H, msg string)  {
	Response(ctx, http.StatusOK,200, data, msg)
}
func Fail(ctx *gin.Context, code int, msg string)  {
	UnifiedResponse(ctx, http.StatusOK,code, msg)
}