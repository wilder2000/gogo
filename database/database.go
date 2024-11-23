package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	"wilder.cn/gogo/config"
	"wilder.cn/gogo/log"
)

var (
	logger   = log.Logger
	DBHander = &gorm.DB{}
)

//	func init() {
//		LoadDatabaseConfig()
//	}
func LoadDatabaseConfig() {

	DBHander = CreateDbHandler(config.AConfig.DataSource)
	printDb(config.AConfig.DataSource)
}
func printDb(db config.DBConfig) {
	logger.DebugF("Init Database%s", db.Name)
	logger.DebugF(db.Name+" %s", db.Type)
	logger.DebugF(db.Name+" %s", db.DSN)
	logger.DebugF(db.Name+"max Idle connections=%d", db.MaxIdleConnections)
	logger.DebugF(db.Name+"max Open connections=%d", db.MaxOpenConnections)
}
func CreateDbHandler(dC config.DBConfig) *gorm.DB {
	myConfig := mysql.Config{
		DSN:               dC.DSN,
		DefaultStringSize: 256,
	}

	sqlLog := glog.Default
	switch log.LConfig.LogLevel {
	case log.LevelDebug:
		sqlLog.LogMode(glog.Info)
		fmt.Println("database log leve: Info")
	case log.LevelInfo:
		sqlLog.LogMode(glog.Info)
		fmt.Println("database log leve: info")
	case log.LevelError:
		sqlLog.LogMode(glog.Error)
		fmt.Println("database log leve: error")
	default:
		sqlLog.LogMode(glog.Silent)
		fmt.Println("database log leve: default silent")
	}

	db, err := gorm.Open(mysql.New(myConfig), &gorm.Config{
		Logger: sqlLog,
	})

	if err != nil {
		panic(err)
	}
	sqldb, err2 := db.DB()
	if err2 != nil {
		logger.ErrorF("Create DB Pool failed %s", err2.Error())
	}
	sqldb.SetMaxOpenConns(dC.MaxOpenConnections)
	sqldb.SetMaxIdleConns(dC.MaxIdleConnections)
	return db
}
func Like(para string) string {
	return "%" + para + "%"
}
