/*************************************************************************
> File Name: adminMiddleware.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-03-10 23:07:47 星期一
> Content: This is a desc
*************************************************************************/

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgs921107/go_framework/common"
	"github.com/sgs921107/go_framework/utils/response"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role.(float64) != float64(common.Administrator) {
			resp := response.NewBaseResponse()
			resp.Error(response.Forbidden)
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}
		c.Next()
	}
}
