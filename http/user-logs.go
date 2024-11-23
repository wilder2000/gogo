package http

import (
	"wilder.cn/gogo/database"
	log2 "wilder.cn/gogo/log"
	"wilder.cn/gogo/sys/db"
)

func dblog(us db.SUser, ip string) {
	log := db.SLog{
		IP:      ip,
		Account: us.ID,
		//Logintime: comm.LocalTime(),
	}
	res := database.DBHander.Create(&log)
	if res.RowsAffected == 1 {
		log2.Logger.InfoF("login log write success.")
	} else {
		log2.Logger.InfoF("login log write failed.")
	}
}
