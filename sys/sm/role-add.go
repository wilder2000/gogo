package sm

import (
	"github.com/gin-gonic/gin"
	hp "net/http"
	"strings"
	"wilder.cn/gogo/comm"
	"wilder.cn/gogo/database"
	"wilder.cn/gogo/http"
	"wilder.cn/gogo/log"
	"wilder.cn/gogo/sys/db"
)

type RoleAddController struct {
	http.AbstractController[AddRoleRequest]
}

func (s RoleAddController) UrlPath() string {
	return "/radd"
}

func (s RoleAddController) Execute(para *AddRoleRequest, c *gin.Context) {
	sr := db.SRole{
		Name:       para.Name,
		Createtime: comm.LocalTime(),
	}
	db := database.DBHander.Create(&sr)
	if db.RowsAffected == 1 {
		c.JSON(hp.StatusOK, http.SuccessResponse("Role添加成功"))
	} else {
		errMsg := db.Error.Error()
		log.Logger.ErrorF("Add role failed.%s", errMsg)
		if strings.Index(errMsg, "1062") > 0 {
			c.JSON(hp.StatusOK, http.FailedResponseCode(http.DataExistFound, "重复的角色", db.Error.Error()))
		} else {
			c.JSON(hp.StatusOK, http.FailedResponse("Role增加失败", db.Error.Error()))
		}
	}
}

type AddRoleRequest struct {
	Name string `json:"name" binding:"required,min=2,max=12"`
}
