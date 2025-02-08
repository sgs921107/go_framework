/*************************************************************************
> File Name: view.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-22 21:33:58 星期五
> Content: This is a desc
*************************************************************************/

package author

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// validator "github.com/go-playground/validator/v10"
	"github.com/sgs921107/go_framework/common"
	"github.com/sgs921107/go_framework/models"
	"github.com/sgs921107/go_framework/utils/paginate"
	"github.com/sgs921107/go_framework/utils/response"
)

var (
	author = models.Author{}
)

// 定义接收数据的结构体
type addForm struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Name     string `form:"name" json:"name" binding:"required,min=1,max=20"`
	Birthday string `form:"birthday" json:"birthday" binding:"required,date"`
}

type addOutput struct {
	response.BaseResponse
}

type getOutput struct {
	response.BaseResponse
	paginate.Paginator[models.AuthorOUtput]
}

// Get 获取所有作者信息
// @BasePath api/v1/author
// @Summary
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 作者相关接口
// @Accept application/json
// @Produce application/json
// @Param object query paginate.PageQuery true "页数"
// @Security ApiKeyAuth
// @Success 200 {object} getOutput "success"
// @Router / [get]
func GetAll(c *gin.Context) {
	var output getOutput
	// 解析参数
	if err := c.BindQuery(&output.Paginator.PageQuery); err != nil {
		output.Error(response.ErrParams, err.Error())
		c.JSON(http.StatusBadRequest, output)
		return
	}
	// 查询数据
	if err := output.Paginator.Paginate(author.Authors()); err != nil {
		output.Error(response.ErrGetAuthors, err.Error())
		c.JSON(http.StatusInternalServerError, output)
		return
	} else {
		output.Ok()
		c.JSON(http.StatusOK, output)
	}
}

// Add 添加一个作者
// @BasePath /api/v1/author
// @Summary 添加一个作者
// @Description 添加一个作者信息 包括名字和生日
// @Tags 作者相关接口
// @Accept application/urlencoding
// @Produce application/json
// @Param data body addForm true "作者生日"
// @Security ApiKeyAuth
// @Success 200 {object} addOutput "Success"
// @Router / [post]
func Add(c *gin.Context) {
	// 声明接收的变量
	var form addForm
	output := addOutput{}
	// Bind()默认解析并绑定form格式
	if err := c.Bind(&form); err != nil {
		output.Error(response.ErrBindForm, err.Error())
		c.JSON(http.StatusBadRequest, output)
		return
	}
	// 校验生日
	birthday, err := time.Parse(common.DATE_LAYOUT, form.Birthday)
	if err != nil {
		output.Error(response.ErrParams, err.Error())
		c.JSON(http.StatusBadRequest, output)
		return
	}
	if err := author.Add(&models.Author{
		Name:     form.Name,
		Birthday: birthday,
	}); err != nil {
		output.Error(response.FailedAddAuthor, err.Error())
		c.JSON(http.StatusInternalServerError, output)
		return
	} else {
		output.Ok()
		c.JSON(http.StatusOK, output)
	}
}
