/*************************************************************************
> File Name: view.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-01-01 20:29:26 星期三
> Content: This is a desc
*************************************************************************/

package users

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/sgs921107/go_framework/app/middlewares"
	"github.com/sgs921107/go_framework/common"
	"github.com/sgs921107/go_framework/models"
	"github.com/sgs921107/go_framework/utils/response"
	"github.com/sgs921107/go_framework/utils/token"
)

type RegisterUser struct {
	Username string `form:"username" json:"username" binding:"required,min=8,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=9,max=30"`
	Role     int    `form:"role" json:"role" binding:"default=0,gte=0,lte=1"`
}

type UpdateUserForm struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=4,max=10"`
	Email    string `form:"Email" json:"Email" binding:"required,email,min=8,max=20"`
	Phone    string `form:"phone" json:"phone" binding:"required,mobile,min=11,max=11"`
	Gender   int    `form:"gender" json:"gender" binding:"required,gte=0,lte=2"`
}

type UserResponse struct {
	response.BaseResponse
	Data *models.UserOutput `json:"data"`
}

// Register 用户注册
// @BasePath /api/v1
// @Summary 添加一个用户
// @Description 用户注册接口
// @Tags 用户相关接口
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param user formData RegisterUser true "注册信息"
// @Success 200 {object} response.BaseResponse "成功"
// @Failure 400 {object} response.BaseResponse "错误的请求"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /users/register [post]
func Register(c *gin.Context) {
	// 声明接收的变量
	var form RegisterUser
	resp := response.NewBaseResponse()
	entry := common.Logger.WithFields(logrus.Fields{
		"ip":  c.ClientIP(),
		"tag": c.Request.URL.Path,
	})
	// 绑定 JSON 请求数据
	if err := c.ShouldBind(&form); err != nil {
		resp.Error(response.ErrParams)
		c.JSON(http.StatusBadRequest, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to bind form")
		return
	}
	// todo
	// 对用户名和密码进行校验
	// 查看用户是否已存在

	entry = entry.WithField("username", form.Username)
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		// 密码加密失败
		resp.Error(response.FailedEncryptPasswd)
		c.JSON(http.StatusInternalServerError, resp)
		entry.WithFields(logrus.Fields{
			"password": form.Password,
			"errMsg":   err.Error(),
		}).Error("Failed to encrypt password")
		return
	}
	user := &models.Users{
		Username: form.Username,
		Password: string(hashedPassword),
		Role:     form.Role,
	}

	// 创建用户
	if err := user.Insert(); err != nil {
		// 插入用户失败
		resp.Error(response.FailedInsertUser)
		c.JSON(http.StatusInternalServerError, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to insert user")
		return
	}
	resp.Ok("用户注册成功")
	c.JSON(http.StatusOK, resp)
}

// Login 用户登录
// @BasePath /api/v1
// @Summary 登录
// @Description 用户登录接口
// @Tags 用户相关接口
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param user formData RegisterUser true "登录信息"
// @Success 200 {object} response.BaseResponse "成功"
// @Failure 400 {object} response.BaseResponse "错误的请求"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /users/login [post]
func Login(c *gin.Context) {
	// 声明接收的变量
	var form RegisterUser
	resp := response.NewBaseResponse()
	entry := common.Logger.WithFields(logrus.Fields{
		"ip":  c.ClientIP(),
		"tag": c.Request.URL.Path,
	})
	// 绑定 JSON 请求数据
	if err := c.ShouldBind(&form); err != nil {
		entry.WithField("errMsg", err.Error()).Error("Failed to bind form")
		resp.Error(response.ErrParams)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// todo
	// 对用户名和密码进行校验
	// 查看用户是否已存在
	entry = entry.WithField("username", form.Username)
	// 查询用户
	user := &models.Users{
		Username: form.Username,
	}
	if err := user.GetByUsername(); err != nil {
		// 用户不存在
		resp.Error(response.FailedAuthorize)
		c.JSON(http.StatusUnauthorized, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to query user by username")
		return
	}
	entry = entry.WithField("userID", user.ID)
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		// 密码错误
		resp.Error(response.FailedAuthorize)
		c.JSON(http.StatusUnauthorized, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to compare password")
		return
	}

	// 生成token
	jwt := token.NewJWT()
	jwtToken, err := jwt.GenWithMap(map[string]interface{}{
		"username": user.Username,
		"user_id":  user.ID,
		"role":     user.Role,
	})
	if err != nil {
		// 生成token失败
		resp.Error(response.FailedGenToken)
		c.JSON(http.StatusInternalServerError, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to generate token")
		return
	}
	// resp header中添加token
	c.Header("X-Auth-Token", jwtToken)

	resp.Ok("登录成功")
	c.JSON(http.StatusOK, resp)
	entry.WithField("userID", user.ID).Info("login success")
}

// UserInfo 用户信息
// @BasePath /api/v1
// @Summary 用户信息
// @Description 查询用户信息接口
// @Tags 用户相关接口
// @Accept application/urlencoded
// @Produce application/json
// @Param id path int true "用户id"
// @Security ApiKeyAuth
// @Success 200 {object} UserResponse "成功"
// @Failure 400 {object} UserResponse "错误的请求"
// @Failure 500 {object} UserResponse "服务器内部错误"
// @Router /users/{id} [get]
func UserInfo(c *gin.Context) {
	resp := UserResponse{}
	entry := common.Logger.WithFields(logrus.Fields{
		"ip":  c.ClientIP(),
		"tag": c.Request.URL.Path,
	})
	// 获取当前用户的user_id
	userID, err := middlewares.GetUserID(c)
	if err != nil {
		resp.Error(response.InvalidUserID)
		c.JSON(http.StatusInternalServerError, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to got user id from token")
		return
	}
	queryID := c.Param("id")
	entry = common.Logger.WithFields(logrus.Fields{
		"userID":  userID,
		"queryID": queryID,
	})
	user := models.Users{}
	if _id, err := common.Str2Uint(queryID); err != nil {
		// query user id不合法
		resp.Error(response.ErrParams)
		c.JSON(http.StatusBadRequest, resp)
		entry.Error("Invalid query ID")
		return
	} else {
		user.ID = _id
	}
	// 查询用户信息
	if err := user.GetByID(); err != nil {
		// 查询用户信息失败
		resp.Error(response.FailedQueryUser)
		c.JSON(http.StatusInternalServerError, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to query user by id")
		return
	}
	// 返回用户信息
	resp.Data = user.Output()
	// 如果查询的不是自己的个人信息则对隐私进行加密
	if userID != user.ID {
		resp.Data.Phone = ""
		resp.Data.Email = ""
	}
	resp.Ok("获取用户信息成功")
	c.JSON(http.StatusOK, resp)
}

// UpdateUserInfo 更新用户信息
// @BasePath /api/v1
// @Summary 更新用户信息
// @Description 更新用户信息接口
// @Tags 用户相关接口
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param user formData UpdateUserForm true "需要更新的用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.BaseResponse "成功"
// @Failure 400 {object} response.BaseResponse "错误的请求"
// @Failure 500 {object} response.BaseResponse "服务器内部错误"
// @Router /users [put]
func UpdateUser(c *gin.Context) {
	// 获取user id
	resp := response.NewBaseResponse()
	entry := common.Logger.WithFields(logrus.Fields{
		"ip":  c.ClientIP(),
		"tag": c.Request.URL.Path,
	})
	// 获取user id
	userID, err := middlewares.GetUserID(c)
	if err != nil {
		resp.Error(response.InvalidUserID)
		c.JSON(http.StatusInternalServerError, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to got user id from token")
		return
	}
	entry.WithField("userID", userID)
	var form UpdateUserForm
	// 绑定 JSON 请求数据
	if err := c.ShouldBind(&form); err != nil {
		resp.Error(response.ErrParams)
		c.JSON(http.StatusBadRequest, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to bind form")
		return
	}

	// todo 校验数据

	entry.WithField("form", form)
	user := &models.Users{}
	user.ID = userID
	if err := user.Update(form); err != nil {
		resp.Error(response.FailedUpdateUser)
		c.JSON(http.StatusInternalServerError, resp)
		entry.WithField("errMsg", err.Error()).Error("Failed to update user")
		return
	}
	resp.Ok("更新用户信息成功")
	c.JSON(http.StatusOK, resp)
}
