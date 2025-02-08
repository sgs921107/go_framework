/*************************************************************************
> File Name: v1.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-20 22:50:48 星期三
> Content: This is a desc
*************************************************************************/

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sgs921107/go_framework/app/v1/author"
	"github.com/sgs921107/go_framework/app/v1/book"
	"github.com/sgs921107/go_framework/app/v1/swagger"
	"github.com/sgs921107/go_framework/app/v1/user"
)

type Group struct {
	Group *gin.RouterGroup
}

func (g *Group) path() string {
	return "/v1"
}

func (g *Group) Register() {
	v1Group := g.Group.Group(g.path())
	userRouter := user.UserRouter{Group: v1Group}
	userRouter.Register()
	bookRouter := book.BookRouter{Group: v1Group}
	bookRouter.Register()
	authorRouter := author.AuthorRouter{Group: v1Group}
	authorRouter.Register()
	swaggerRouter := swagger.SwaggerRouter{Group: v1Group, Version: "v1"}
	swaggerRouter.Register()
}
