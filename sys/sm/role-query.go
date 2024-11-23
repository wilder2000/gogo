package sm

import (
	"github.com/gin-gonic/gin"
	"wilder.cn/gogo/http"
	"wilder.cn/gogo/sys/db"
)

type RoleQueryController struct {
	http.AbstractController[http.QueryRequest]
}

func (s RoleQueryController) UrlPath() string {
	return "/rquery"
}

func (s RoleQueryController) Execute(para *http.QueryRequest, c *gin.Context) {
	http.QueryPage[db.SRole](para, c)
}
