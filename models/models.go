/*************************************************************************
> File Name: models.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-21 10:52:26 星期四
> Content: This is a desc
*************************************************************************/

package models

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	mysqlClient := MysqlClient{}
	db = mysqlClient.DB()
	// 自动迁移
	mysqlClient.Migrator()
}
