package sm

import (
	"github.com/gin-gonic/gin"
	hp "net/http"
	"wilder.cn/gogo/http"
)

type PwdController struct {
	http.AbstractController[http.ChangePWD]
}

func (s PwdController) UrlPath() string {
	return "/pwd"
}

func (s PwdController) Execute(para *http.ChangePWD, c *gin.Context) {
	err := http.UserProxy.ChangePwd(para.Password, para.Email)
	if err != nil {
		c.JSON(hp.StatusOK, http.FailedResponse("update user password failed", err))
	} else {
		c.JSON(hp.StatusOK, http.SuccessResponse("update user password success"))
	}

}
