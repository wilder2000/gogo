package sm

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"math"
	hp "net/http"
	"strconv"
	"strings"
	"wilder.cn/gogo/comm"
	"wilder.cn/gogo/database"
	"wilder.cn/gogo/http"
	"wilder.cn/gogo/log"
	"wilder.cn/gogo/sys/db"
)

const (
	DepUserSql = "select u.id,u.email,du.departmentid,departmentid is not null as selected from s_users u left join s_depusers du  on u.id=du.userid and du.departmentid=@did where u.email like '%@uemail%' order by selected desc"
	Did        = "did"
)

type DepartmentController struct {
	http.AbstractController[http.QueryRequest]
}

func (s DepartmentController) UrlPath() string {
	return "/dm"
}
func (s DepartmentController) Execute(req *http.QueryRequest, c *gin.Context) {
	log.Logger.InfoF("Received query:%s", c.Request.RequestURI)
	switch req.Code {

	case http.CmdQueryUnionDepartments:
		//查询部门拥有的用户，并上没有加入的用户
		rsql := strings.ReplaceAll(DepUserSql, "@uemail", comm.IToString(req.Where["name"]))
		findDepUsers[db.SUserWithDep](req, c, rsql, Did, Did)

	default:
		c.JSON(hp.StatusOK, http.FailedResponse("Not defined cmd:"+strconv.Itoa(req.Code), ""))
	}
}

func findDepUsers[T any](req *http.QueryRequest, c *gin.Context, rawSql string, key string, sqlName string) {
	name, ok := req.Where[key]
	if !ok {
		c.JSON(hp.StatusOK, http.FailedResponse("Not found '"+key+"' in where map", ""))
		return
	}
	page := http.RequestToPage[T](*req)
	var rowData []*T
	did := comm.IToString(name)
	log.Logger.InfoF("did%s", did)
	db := database.DBHander.Raw(rawSql, sql.Named(sqlName, did))
	var totalRows int64
	db.Debug().Count(&totalRows)
	log.Logger.InfoF("count totalRows=%d", totalRows)
	log.Logger.InfoF("count error=", db.Error)

	page.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(page.PageSize)))
	page.TotalPages = totalPages

	log.Logger.InfoF("limit:%d", page.Limit())
	log.Logger.InfoF("Offset:%d", page.Offset())
	rawSql += " limit @limit offset @offset"
	db2 := database.DBHander.Raw(rawSql, sql.Named(sqlName, did), sql.Named("limit", page.Limit()), sql.Named("offset", page.Offset()))

	db2.Debug().Find(&rowData)
	page.Rows = rowData
	c.JSON(hp.StatusOK, http.SuccessResponse(http.PageToResponse(page)))
}
