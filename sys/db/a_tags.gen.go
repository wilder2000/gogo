// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"gorm.io/gorm"
	"time"
	"wilder.cn/gogo/comm"
)

const TableNameATag = "a_tags"

// ATag mapped from table <a_tags>
type ATag struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`                      // 标签名称
	Createtime time.Time `gorm:"column:createtime;not null" json:"createtime"` // 创建时间
	Sumdocs    int32     `gorm:"column:sumdocs" json:"sumdocs"`                // 文档数量
}

// TableName ATag's table name
func (*ATag) TableName() string {
	return TableNameATag
}
func (ad *ATag) BeforeCreate(tx *gorm.DB) error {
	ad.Createtime=comm.LocalTime()
	return nil
}