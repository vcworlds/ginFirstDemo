package middleware

import (
	"gin-Vue/dao"
	"gin-Vue/models"
	"gin-Vue/response"
	"gin-Vue/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// 对token进行判断
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Error(ctx, "token无效")
			ctx.Abort()
			return
		}
		// 解析token
		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Error(ctx, "token已过期")
			ctx.Abort()
			return
		}
		var user models.User
		dao.DB.Take(&user, claims.Id)
		if user.ID == 0 {
			response.Error(ctx, "该用户开了小差")
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
