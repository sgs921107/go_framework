/*************************************************************************
> File Name: user.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-01-01 20:28:47 星期三
> Content: This is a desc
*************************************************************************/

package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Group *gin.RouterGroup
}

func (r *UserRouter) path() string {
	return "/user"
}

// Register register user router group
// @BasePath /api/v1/user
func (r *UserRouter) Register() {
	router := r.Group.Group(r.path())
	router.GET("/register", Register)
}
