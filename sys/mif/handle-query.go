package mif

import (
	"github.com/gin-gonic/gin"
	gh "net/http"
	"wilder.cn/gogo/http"
)

const REQQuery = "/mif/q"

func HandleQuery(c *gin.Context) {
	if ValidHttpMethods(c) {
		var paraModel QRequest
		if err := c.ShouldBind(&paraModel); err == nil {
			//next
			hf, exist := TargetQueryFunc[paraModel.Target]
			if !exist {
				c.JSON(gh.StatusOK, http.FailedResponse("Not found mif data map target %"+paraModel.Target, ""))
				return
			}
			hf(&paraModel, c)
			return
		} else {
			http.HandlErr(c, err)
		}
	}

}
