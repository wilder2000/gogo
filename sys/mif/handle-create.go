package mif

import (
	"github.com/gin-gonic/gin"
	gh "net/http"
	"wilder.cn/gogo/http"
)

const REQCreate = "/mif/c"

func HandleCreate(c *gin.Context) {
	if ValidHttpMethods(c) {
		var paraModel ARequest
		if err := c.ShouldBind(&paraModel); err == nil {
			hf, exist := TargetCreateFunc[paraModel.Target]
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
