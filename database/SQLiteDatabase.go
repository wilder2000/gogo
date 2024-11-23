package database

//import (
//	"gorm.io/gorm"
//	"wilder.cn/gogo/log"
//)
//import "gorm.io/driver/sqlite"
//
//type SQLiteHandler struct {
//	DbFile string
//}
//
//func (r SQLiteHandler) Open() (*gorm.DB, bool) {
//	db, err := gorm.Open(sqlite.Open(r.DbFile), &gorm.Config{})
//	if err != nil {
//		log.Logger.ErrorF("Try to open sqlite failed.%s", err.Error())
//
//		return nil, false
//	}
//	return db, true
//}
