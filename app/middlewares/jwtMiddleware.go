/*************************************************************************
> File Name: jwtMiddleware.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-01-02 14:40:43 星期四
> Content: This is a desc
*************************************************************************/

package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/sgs921107/go_framework/common"
	"github.com/sgs921107/go_framework/utils/response"
	"github.com/sgs921107/go_framework/utils/token"
	"github.com/sirupsen/logrus"
)

// JwtMiddleware json web token 中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token
		bearerToken := ctx.Request.Header.Get("Authorization")
		entry := common.Logger.WithFields(logrus.Fields{
			"ip":    ctx.ClientIP(),
			"tag":   "jwtMiddleware",
			"token": bearerToken,
		})
		resp := response.BaseResponse{}
		if bearerToken == "" {
			resp.Error(response.Unauthorized)
			ctx.JSON(http.StatusUnauthorized, resp)
			ctx.Abort()
			entry.Error("Empty token")
			return
		}
		tokenStr := strings.TrimPrefix(bearerToken, "Bearer ")
		// 校验token
		jwt := token.NewJWT()
		claims, err := jwt.Parse(tokenStr)
		switch err {
		case nil:
			ctx.Set("user_id", claims["user_id"])
			ctx.Next()
		// token已过期
		case gojwt.ErrTokenExpired:
			// 判断是否可刷新token
			ok, err := jwt.AllowRefresh(claims)
			if ok {
				// 重新生成token
				if newToken, err := jwt.GenWithClaims(claims); err == nil {
					resp.Error(response.ErrTokenExpried1)
					// resp header中添加token
					ctx.Header("X-Auth-Token", newToken)
					entry.WithField("newToken", newToken).Info("Refershed token")
				} else {
					resp.Error(response.ErrTokenExpried2)
					entry.WithField("errMsg", err.Error()).Info("Failed to refersh token")
				}
			} else {
				resp.Error(response.ErrTokenExpried2)
			}
			ctx.JSON(http.StatusUnauthorized, resp)
			ctx.Abort()
			entry.WithField("errMsg", err.Error()).Error("Expried Token")
		default:
			resp.Error(response.InvalidToken)
			ctx.JSON(http.StatusUnauthorized, resp)
			ctx.Abort()
			entry.WithField("errMsg", err.Error()).Error("Failed to parse token")
		}
	}
}

// 从请求上下文中获取user id type: uint
func GetUserID(c *gin.Context) (uint, error) {
	// 获取当前用户的user_id
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("Failed to get user id from token")
	}
	float6464UserID, ok := userID.(float64)
	if !ok || float6464UserID < 0 {
		return 0, errors.New("Got invalid user id from token")
	}
	return uint(float6464UserID), nil
}
