/*************************************************************************
> File Name: view.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-20 23:06:34 星期三
> Content: This is a desc
*************************************************************************/

package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgs921107/go_framework/models"
)

var (
	book = models.Book{}
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"books": book.Books(),
	})
}
