package mif

import (
	"github.com/gin-gonic/gin"
	gh "net/http"
	"wilder.cn/gogo/http"
)

const REQDelete = "/mif/d"

func HandleDelete(c *gin.Context) {
	if ValidHttpMethods(c) {
		var paraModel DRequest
		if err := c.ShouldBind(&paraModel); err == nil {
			//next
			hf, exist := TargetDeleteFunc[paraModel.Target]
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
