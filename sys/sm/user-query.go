package sm

import (
	"github.com/gin-gonic/gin"
	hp "net/http"
	"wilder.cn/gogo/http"
	"wilder.cn/gogo/log"
	"wilder.cn/gogo/sys/db"
)

type UserQueryController struct {
	http.AbstractController[http.QueryRequest]
}

func (s UserQueryController) UrlPath() string {
	return "/uquery"
}

func (s UserQueryController) Execute(para *http.QueryRequest, c *gin.Context) {
	log.Logger.InfoF("Received query:%s", c.Request.RequestURI)
	con := para.Where
	log.Logger.InfoF("query map:%v", *para)

	pg := http.NewPage[db.SUser]()
	pg.PageSize = para.PageSize
	pg.PageIndex = para.PageIndex
	res, err := http.SelectPage(con, pg)
	if err != nil {
		c.JSON(hp.StatusOK, http.FailedResponse("query failed", err))
	} else {
		qres := &http.QueryResponse[db.SUser]{}
		qres.PageSize = res.PageSize
		qres.PageIndex = res.PageIndex
		qres.TotalPages = res.TotalPages
		qres.Data = res.Rows
		c.JSON(hp.StatusOK, http.SuccessResponse(qres))
	}
}
