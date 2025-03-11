/*************************************************************************
> File Name: user.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-01-01 20:28:47 星期三
> Content: This is a desc
*************************************************************************/

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sgs921107/go_framework/app/middlewares"
)

type UsersRouter struct {
	Group *gin.RouterGroup
}

func (r *UsersRouter) path() string {
	return "/users"
}

// Register register user router group
// @BasePath /api/v1/users
func (r *UsersRouter) Register() {
	router := r.Group.Group(r.path())
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.GET("/:id", middlewares.JwtMiddleware(), UserInfo)
	router.PUT("", middlewares.JwtMiddleware(), UpdateUser)
}
