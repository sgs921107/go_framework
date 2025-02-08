/*************************************************************************
> File Name: author.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-22 21:32:40 星期五
> Content: This is a desc
*************************************************************************/

package author

import (
	"github.com/gin-gonic/gin"
)

type AuthorRouter struct {
	Group *gin.RouterGroup
}

func (r *AuthorRouter) path() string {
	return "/author"
}

func (r *AuthorRouter) Register() {
	// @BasePath /api/v1/author
	router := r.Group.Group(r.path())
	router.GET("/", GetAll)
	router.POST("/", Add)
}
