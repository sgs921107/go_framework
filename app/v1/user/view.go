/*************************************************************************
> File Name: view.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-01-01 20:29:26 星期三
> Content: This is a desc
*************************************************************************/

package user

import (
	"github.com/gin-gonic/gin"
)

type RegisterUser struct {
	Username string `form:"username" json:"username" binding:"required,min=6,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=9,max=20"`
}

// Register 用户注册
// @BasePath /api/v1/user
// @Summary 添加一个用户
// @Description 用户注册接口
// @Tags 用户相关接口
// @Accept application/urlencoding
// @Produce application/json
// @Param data body addForm true "用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} addOutput "Success"
// @Router /register [post]
func Register(c *gin.Context) {
	// 声明接收的变量
	// var user RegisterUser
}
