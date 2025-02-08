/*************************************************************************
> File Name: jwtMiddleware.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-01-02 14:40:43 星期四
> Content: This is a desc
*************************************************************************/

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/sgs921107/go_framework/utils/response"
	"github.com/sgs921107/go_framework/utils/token"
)

// JwtMiddleware json web token 中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token
		tokenStr := ctx.Request.Header.Get("Authotization")
		output := response.BaseResponse{}
		if tokenStr == "" {
			output.Error(response.NoUnauthorized, "用户未登录")
			ctx.JSON(http.StatusUnauthorized, output)
			return
		}
		// 校验token
		jwt := token.NewJWT()
		switch jwt.Valid(tokenStr) {
		case nil:
			ctx.Next()
		// token已过期
		case gojwt.ErrTokenExpired:
			// 判断是否可刷新token
			if ok, err := jwt.AllowRefresh(tokenStr); ok {
				output.Error(response.ErrTokenExpried1, "token已过期, 请刷新token")
			} else {
				var errMsg string
				if err == nil {
					errMsg = "token已过期吗, 请重新登录"
				} else {
					errMsg = err.Error()
				}
				output.Error(response.ErrTokenExpried2, errMsg)
			}
			ctx.JSON(http.StatusUnauthorized, output)
		default:
			output.Error(response.ErrTokenInvalid, "token验证失败")
			ctx.JSON(http.StatusUnauthorized, output)
		}
	}
}
