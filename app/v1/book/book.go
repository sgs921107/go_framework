/*************************************************************************
> File Name: book.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-20 23:05:01 星期三
> Content: This is a desc
*************************************************************************/

package book

import (
	"github.com/gin-gonic/gin"
)

type BookRouter struct {
	Group *gin.RouterGroup
}

func (r *BookRouter) path() string {
	return "/book"
}

// Register register book router group
// @BasePath /api/v1/book
func (r *BookRouter) Register() {
	router := r.Group.Group(r.path())
	router.GET("/", Get)
}
