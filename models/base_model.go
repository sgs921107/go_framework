/*************************************************************************
> File Name: base_model.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-22 19:24:47 星期五
> Content: This is a desc
*************************************************************************/

package models

import (
	"github.com/bwmarrin/snowflake"
	"github.com/sgs921107/go_framework/common"
	"gorm.io/gorm"
)

var (
	snowFlakeNode *snowflake.Node
)

// 雪花model 使用雪花算法来生成id
type SnowFlakeModel struct {
	gorm.Model
}

func (m *SnowFlakeModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uint(NextId())
	return
}

// 获取id
func NextId() int64 {
	return int64(snowFlakeNode.Generate())
}

// 初始化雪花ID生成器
func init() {
	var err error
	snowFlakeNode, err = snowflake.NewNode(int64(common.Setting.Snow.NodeID))
	if err != nil {
		common.Logger.WithField("err", err.Error()).Fatal("Failed To New Snow Flake Node!")
	}
}
