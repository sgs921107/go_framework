/*************************************************************************
> File Name: book.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-21 20:27:15 星期四
> Content: This is a desc
*************************************************************************/

package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name            string    `gorm:"index;column:name;size:50;not null;comment:书名"`
	PublicationDate time.Time `gorm:"column:publication_date;type:date;not null;comment:出版日期"`
	// 外键
	Author   Author `gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID"`
	AuthorID int    `gorm:"column:author_id;comment:外键: 作者id"`
}

func (b *Book) Books() (books []Book) {
	db.Model(b).Select("book.name", "book.publication_date", "author.name").Joins("right join author on book.author_id=author.id").Find(&books)
	return books
}
