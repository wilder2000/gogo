package sm

import (
	"github.com/gin-gonic/gin"
	hp "net/http"
	"wilder.cn/gogo/http"
)

type CheckPwdController struct {
	http.AbstractController[http.CheckPWD]
}

func (s CheckPwdController) UrlPath() string {
	return "/cpwd"
}

func (s CheckPwdController) Execute(para *http.CheckPWD, c *gin.Context) {
	lUser := &http.LoginInput{
		Email:    para.Email,
		Password: para.Password,
	}
	_, err2 := http.UserProxy.Login(*lUser)
	if err2 != nil {
		c.JSON(hp.StatusOK, http.FailedResponseCode(err2.Code(), "pwd is invalid", err2.Error()))
		return
	}
	c.JSON(hp.StatusOK, http.SuccessResponse(para))
}
