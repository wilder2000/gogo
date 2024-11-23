package sm

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	hp "net/http"
	"strconv"
	"wilder.cn/gogo/database"
	"wilder.cn/gogo/http"
	"wilder.cn/gogo/sys/db"
)

type UserProfileController struct {
	http.AbstractController[http.GetRequest]
}

func (s UserProfileController) UrlPath() string {
	return "/upro"
}

func (s UserProfileController) Execute(para *http.GetRequest, c *gin.Context) {
	cmd := para.Code
	switch cmd {
	case http.CmdGetUserBasic:
		fn := para.FilterName
		fv := para.FilterValue
		var usr db.SUser
		db := database.DBHander.Raw("select * from s_users where id=@uid", sql.Named(fn, fv)).Find(&usr)
		if db.Error != nil {
			c.JSON(hp.StatusOK, http.FailedResponse(db.Error.Error(), ""))
		} else {
			deps, err := http.FindUserDeps(fv)
			if err != nil {
				c.JSON(hp.StatusOK, http.FailedResponse("Load user department failed."+err.Error(), ""))
			} else {
				formatter := http.FormatUser(usr)
				formatter.Department = deps
				c.JSON(hp.StatusOK, http.SuccessResponse(formatter))
			}

		}

	default:
		c.JSON(hp.StatusOK, http.FailedResponse("Not implement cmd:"+strconv.Itoa(cmd), ""))
	}
}
