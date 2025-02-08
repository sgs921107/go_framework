/*************************************************************************
> File Name: authors.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-21 20:08:23 星期四
> Content: This is a desc
*************************************************************************/

package models

import (
	"time"

	"github.com/sgs921107/go_framework/common"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name     string    `gorm:"unique;column:name;size:10;not null;comment:作者名称"`
	Birthday time.Time `gorm:"column:birthday;type:date;not null;comment:生日"`
}

type AuthorOUtput struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
}

func (a *Author) Add(author *Author) error {
	entry := common.Logger.WithField("authors", author)
	result := db.Create(author)
	if result.Error != nil {
		entry.WithField("err", result.Error.Error()).Error("Failed To Create Author!")
		return result.Error
	}
	entry.Info("Succeed To Create Author!")
	return nil
}

func (a *Author) Authors() *gorm.DB {
	return db.Model(a)
}
