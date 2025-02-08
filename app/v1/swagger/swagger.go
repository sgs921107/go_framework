/*************************************************************************
> File Name: swagger.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-22 22:48:16 星期五
> Content: This is a desc
*************************************************************************/

package swagger

import (
	"github.com/gin-gonic/gin"
	docs "github.com/sgs921107/go_framework/app/v1/swagger/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerRouter struct {
	Group   *gin.RouterGroup
	Version string
}

func (r *SwaggerRouter) path() string {
	return "swagger"
}

func (r *SwaggerRouter) Register() {
	docs.SwaggerInfo.Version = r.Version
	docs.SwaggerInfo.BasePath = r.Group.BasePath()
	group := r.Group.Group(r.path())
	group.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
