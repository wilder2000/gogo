package sm

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	GroupSql        = "select * from s_group g left join s_groupuser gu  on g.id=gu.groupid where gu.userid=@uid"
	GroupTrueSql    = "select *,userid is not null as selected from s_group g left join s_groupuser gu  on g.id=gu.groupid and gu.userid=@uid where g.name like '%@gname%' order by selected desc"
	GroupUsersSql   = "select u.id,u.email,gu.groupid,groupid is not null as selected from s_users u left join s_groupuser gu  on u.id=gu.userid and gu.groupid=@gid where u.email like '%@uname%' order by selected desc"
	GroupRolesSql   = "select r.id,r.name,rg.groupid,groupid is not null as selected from s_role r left join s_rolegroup rg  on r.id=rg.roleid and rg.groupid=@gid where r.name like '%@rname%' order by selected desc"
	GroupAllSql     = "select * from s_group where name=@name"
	GroupALLUserSQL = "select g.id,g.name,gu.userid from s_group g left join s_groupuser gu  on g.id=gu.groupid where g.id = ?"
	Uid             = "uid"
	Gid             = "gid"
	Gname           = "name"
)

type UserGroupsController struct {
	http.AbstractController[http.QueryRequest]
}

func (s UserGroupsController) UrlPath() string {
	return "/ug"
}

func (s UserGroupsController) Execute(req *http.QueryRequest, c *gin.Context) {
	log.Logger.InfoF("Received query:%s", c.Request.RequestURI)
	switch req.Code {
	case http.CmdUserGroups:
		findGroup[db.SGroup](req, c, GroupSql, Uid, Uid)
	case http.CmdQueryGroups:
		findGroup[db.SGroup](req, c, GroupAllSql, Gname, Gname)
	case http.CmdQueryUnionRoles:
		//查询编组加入的角色，并上没有加入的角色
		rsql := strings.ReplaceAll(GroupRolesSql, "@rname", comm.IToString(req.Where["name"]))
		findGroup[db.SRoleWithGroup](req, c, rsql, Gid, Gid)
	case http.CmdQueryUnionUsers:
		//查询编组下面的用户，并上没有加入的用户
		usql := strings.ReplaceAll(GroupUsersSql, "@uname", comm.IToString(req.Where["name"]))
		findGroup[db.SUserWithGroup](req, c, usql, Gid, Gid)
	case http.CmdQueryUnionGroups:
		//查询用户加入的编组，加上没有加入的编组，加入的编组，字段selected=true
		gsql := strings.ReplaceAll(GroupTrueSql, "@gname", comm.IToString(req.Where["name"]))
		findGroup[db.SGroupWithUser](req, c, gsql, Uid, Uid)
	default:
		c.JSON(hp.StatusOK, http.FailedResponse("Not defined cmd:"+strconv.Itoa(req.Code), ""))
	}

}

func findGroup[T any](req *http.QueryRequest, c *gin.Context, rawSql string, key string, sqlName string) {
	name, ok := req.Where[key]
	if !ok {
		c.JSON(hp.StatusOK, http.FailedResponse("Not found '"+key+"' in where map", ""))
		return
	}
	page := http.RequestToPage[T](*req)
	var rowData []*T
	gid := comm.IToString(name)
	log.Logger.InfoF("gid%s", gid)
	db := database.DBHander.Raw(rawSql, sql.Named(sqlName, gid))
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
	db2 := database.DBHander.Raw(rawSql, sql.Named(sqlName, gid), sql.Named("limit", page.Limit()), sql.Named("offset", page.Offset()))

	db2.Debug().Find(&rowData)
	page.Rows = rowData
	c.JSON(hp.StatusOK, http.SuccessResponse(http.PageToResponse(page)))
}

// FindUserByGroupID 查询用户组下面所有用户
func FindUserByGroupID(groupid int32) (rows gorm.Rows, err error) {
	rows, err = database.DBHander.Raw(GroupALLUserSQL, groupid).Rows()
	return rows, err
}
