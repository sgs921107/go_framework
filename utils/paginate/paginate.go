/*************************************************************************
> File Name: paginate.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-29 19:36:20 星期五
> Content: This is a desc
*************************************************************************/

package paginate

import (
	"gorm.io/gorm"
)

type PageQuery struct {
	Page int `json:"page" form:"page" binding:"required" default:"1"`  // 当前页
	Size int `json:"size" form:"size" binding:"required" default:"10"` // 当前页的数量
}

type Paginator[T any] struct {
	PageQuery
	Pages   int   `json:"pages"`   // 总页数
	Total   int64 `json:"total"`   // 总数
	HasPerv bool  `json:"hasPerv"` //是否有上一页
	HasNext bool  `json:"hasNext"` // 是否有下一页
	Data    []T   `json:"data"`    // 数据
}

// total 获取总条数
func (p *Paginator[T]) total(stmt *gorm.DB) error {
	return stmt.Count(&p.Total).Error
}

// pages 总页数
func (p *Paginator[T]) pages() int {
	quotient := int(p.Total) / p.Size
	remainder := int(p.Total) % p.Size
	if remainder > 0 {
		quotient++
	}
	return quotient
}

// hasPerv 是否有上一页
func (p *Paginator[T]) hasPerv() bool {
	return p.Page > 1
}

// hasNext 是否有下一页
func (p *Paginator[T]) hasNext() bool {
	return p.Page < p.Pages
}

// data 获取当前页的数据
func (p *Paginator[T]) data(stmt *gorm.DB) error {
	err := stmt.Limit(p.Size).Offset(p.Size * (p.Page - 1)).Find(&p.Data).Error
	return err
}

// 填充分页数据
func (p *Paginator[T]) Paginate(stmt *gorm.DB) error {
	if err := p.total(stmt); err != nil {
		return err
	}
	if err := p.data(stmt); err != nil {
		return err
	}
	p.Pages = p.pages()
	p.HasPerv = p.hasPerv()
	p.HasNext = p.hasNext()
	return nil
}
