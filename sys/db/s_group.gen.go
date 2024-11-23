// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"gorm.io/gorm"
	"time"
	"wilder.cn/gogo/comm"
)

const TableNameSGroup = "s_group"

// SGroup mapped from table <s_group>
type SGroup struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`                      // 组名
	Createtime time.Time `gorm:"column:createtime;not null" json:"createtime"` // 创建时间
}

// TableName SGroup's table name
func (*SGroup) TableName() string {
	return TableNameSGroup
}
func (ad *SGroup) BeforeCreate(tx *gorm.DB) error {
	ad.Createtime=comm.LocalTime()
	return nil
}