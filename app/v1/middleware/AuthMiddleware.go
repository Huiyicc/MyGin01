package middleware

import (
	"gin01/app/v1/common"
	"gin01/app/v1/model"
	"gin01/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")																//读取Token Header头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {										// Token头 Bearer
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":401.1,"msg":"权限不足" + tokenString})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]																					// Bearer + 空格 站7位

		token, claims, err := common.ParseToken(tokenString)															// claims => Token解析后数据
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401.2,"msg":"权限不足2"})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		DB := config.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401.3,"msg":"权限不足","user":user})
			ctx.Abort()
			return
		}
		//用户存在 将user 的信息写入上下文
		ctx.Set("user",user)
		ctx.Next()
	}
}
