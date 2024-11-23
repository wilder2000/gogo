package sm

import (
	"github.com/gin-gonic/gin"
	hp "net/http"
	"strconv"
	"wilder.cn/gogo/database"
	"wilder.cn/gogo/http"
	"wilder.cn/gogo/log"
	"wilder.cn/gogo/sys/db"
)

type UserMgrController struct {
	http.AbstractController[http.UpdateRequest]
}

func (s UserMgrController) UrlPath() string {
	return "/um"
}
func (s UserMgrController) Execute(req *http.UpdateRequest, c *gin.Context) {
	log.Logger.InfoF("Received update:%s", c.Request.RequestURI)
	switch req.Code {
	case http.CmdUpdateUserName:
		uname, okName := req.Fields["name"]
		uid, okID := req.Where["id"]
		if okName && okID {
			db := database.DBHander.Model(db.SUser{}).Where("id=?", uid).Update("name", uname)
			if db.RowsAffected == 1 {
				c.JSON(hp.StatusOK, http.SuccessResponse(""))
			} else {
				c.JSON(hp.StatusOK, http.FailedResponse("Failed to update", db.Error))
			}
		} else {
			c.JSON(hp.StatusOK, http.FailedResponse("Not found the name or id field in request map", ""))
		}
	case http.CmdUpdateUserSex:
		usex, okSex := req.Fields["sex"]
		uid, okUid := req.Where["id"]
		if okSex && okUid {
			db := database.DBHander.Model(db.SUser{}).Where("id=?", uid).Update("sex", usex)
			if db.RowsAffected == 1 {
				c.JSON(hp.StatusOK, http.SuccessResponse(""))
			} else {
				c.JSON(hp.StatusOK, http.FailedResponse("Failed to update", db.Error))
			}
		} else {
			c.JSON(hp.StatusOK, http.FailedResponse("Not found the name or id field in request map", ""))
		}
	default:
		c.JSON(hp.StatusOK, http.FailedResponse("Not defined cmd:"+strconv.Itoa(req.Code), ""))
	}
}
