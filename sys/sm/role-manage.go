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
	RolesOperatorsSql = "select o.id,o.name,ro.roleid,roleid is not null as selected from s_operators o left join s_roleoperator ro  on o.id=ro.operatorid and ro.roleid=@rid where o.name like '%@oname%' order by selected desc"
	Rid               = "rid"
)

type OperatorController struct {
	http.AbstractController[http.QueryRequest]
}

func (s OperatorController) UrlPath() string {
	return "/rm"
}
func (s OperatorController) Execute(req *http.QueryRequest, c *gin.Context) {
	log.Logger.InfoF("Received query:%s", c.Request.RequestURI)
	switch req.Code {

	case http.CmdQueryUnionOperators:
		//查询角色拥有的操作权限，并上没有的操作权限
		rsql := strings.ReplaceAll(RolesOperatorsSql, "@oname", comm.IToString(req.Where["name"]))
		findOperators[db.SOperatorWithRole](req, c, rsql, Rid, Rid)

	default:
		c.JSON(hp.StatusOK, http.FailedResponse("Not defined cmd:"+strconv.Itoa(req.Code), ""))
	}
}

func findOperators[T any](req *http.QueryRequest, c *gin.Context, rawSql string, key string, sqlName string) {
	name, ok := req.Where[key]
	if !ok {
		c.JSON(hp.StatusOK, http.FailedResponse("Not found '"+key+"' in where map", ""))
		return
	}
	page := http.RequestToPage[T](*req)
	var rowData []*T
	rid := comm.IToString(name)
	log.Logger.InfoF("rid%s", rid)
	db := database.DBHander.Raw(rawSql, sql.Named(sqlName, rid))
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
	db2 := database.DBHander.Raw(rawSql, sql.Named(sqlName, rid), sql.Named("limit", page.Limit()), sql.Named("offset", page.Offset()))

	db2.Debug().Find(&rowData)
	page.Rows = rowData
	c.JSON(hp.StatusOK, http.SuccessResponse(http.PageToResponse(page)))
}

//insert into s_urlmappings (operatorid,url) values(10,'/pwd');
//insert into s_urlmappings (operatorid,url) values(10,'/cpwd');
//insert into s_urlmappings (operatorid,url) values(10,'/uquery');
//insert into s_urlmappings (operatorid,url) values(10,'/rquery');
//insert into s_urlmappings (operatorid,url) values(10,'/radd');
//insert into s_urlmappings (operatorid,url) values(10,'/ug');
//insert into s_urlmappings (operatorid,url) values(10,'/rm');
//insert into s_urlmappings (operatorid,url) values(10,'/dm');
//insert into s_urlmappings (operatorid,url) values(10,'/um');
//insert into s_urlmappings (operatorid,url) values(10,'/upro');
//insert into s_urlmappings (operatorid,url) values(10,'/mif/c');
//insert into s_urlmappings (operatorid,url) values(10,'/mif/q');
//insert into s_urlmappings (operatorid,url) values(10,'/mif/d');
//insert into s_urlmappings (operatorid,url) values(10,'/mif/u');
