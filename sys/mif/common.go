package mif

import (
	"github.com/gin-gonic/gin"
	"wilder.cn/gogo/sys/db"
)

//type QRequest struct {
//	PageIndex int            `json:"PageIndex"`
//	PageSize  int            `json:"PageSize"`
//	Target    string         `json:"Target"` //目标对象
//	Where     map[string]any `json:"Where"`
//}
//type ARequest struct {
//	Target string         `json:"Target"` //目标对象名称
//	Fields map[string]any `json:"Fields"`
//}
//
//type URequest struct {
//	Target string         `json:"Target"` //目标对象名称
//	Fields map[string]any `json:"Fields"`
//	Where  map[string]any `json:"Where"`
//}
//
//type DRequest struct {
//	Target string         `json:"Target"` //目标对象名称
//	Where  map[string]any `json:"Where"`
//}
//type QResponse struct {
//	PageIndex  int    `json:"PageIndex"`
//	PageSize   int    `json:"PageSize"`
//	TotalPages int    `json:"TotalPages"`
//	Message    string `json:"message"`
//	Code       int    `json:"code"`
//	Data       any    `json:"Data"`
//}

var (
	TargetQueryFunc  = make(map[string]HandleQueryTarget)
	TargetCreateFunc = make(map[string]HandleCreateTarget)
	TargetDeleteFunc = make(map[string]HandleDeleteTarget)
	TargetUpdateFunc = make(map[string]HandleUpdateTarget)

	TargetPreUpdateFunc = make(map[string]func())
)

//type HandleQueryTarget func(para *QRequest, c *gin.Context)
//type HandleCreateTarget func(para *ARequest, c *gin.Context)
//type HandleDeleteTarget func(para *DRequest, c *gin.Context)
//type HandleUpdateTarget func(para *URequest, c *gin.Context)

func init() {
	RegObject[db.SUser]("user")
	RegObject[db.SRole]("role")
	RegObject[db.SGroup]("group")
	RegObject[db.SGroupuser]("groupuser")
	RegObject[db.SRolegroup]("rolegroup")
	RegObject[db.SOperator]("operator")
	RegObject[db.SRoleoperator]("roleoper")
	RegObject[db.SDepartment]("depart")
	RegObject[db.SDepuser]("depuser")

	RegObject[db.ATag]("tags")
}
func RegObject[T any](t string) {
	TargetQueryFunc[t] = func(para *QRequest, c *gin.Context) { QueryObject[T](para, c) }
	TargetCreateFunc[t] = func(para *ARequest, c *gin.Context) { CreateObject[T](para, c) }
	TargetDeleteFunc[t] = func(para *DRequest, c *gin.Context) { DeleteObject[T](para, c) }
	TargetUpdateFunc[t] = func(para *URequest, c *gin.Context) { UpdateObject[T](para, c) }
}

//
//const (
//	DefPageSize = 20
//)
//
//type Page[T any] struct {
//	PageSize   int    `json:"pageSize,omitempty" form:"pageSize"`
//	PageIndex  int    `json:"pageIndex,omitempty" form:"pageIndex"`
//	Sort       string `json:"sort,omitempty" form:"sort"`
//	TotalRows  int64  `json:"total_rows"`
//	TotalPages int    `json:"total_pages"`
//	Rows       []T    `json:"rows"`
//}
//
//func (p *Page[T]) Offset() int {
//	return (p.CurrentPage() - 1) * p.Limit()
//}
//
//func (p *Page[T]) Limit() int {
//	if p.PageSize == 0 {
//		p.PageSize = DefPageSize
//	}
//	return p.PageSize
//}
//
//func (p *Page[T]) CurrentPage() int {
//	if p.PageIndex == 0 {
//		p.PageIndex = 1
//	}
//	return p.PageIndex
//}
//
//func (p *Page[T]) GetSort() string {
//	if p.Sort == "" {
//		p.Sort = "Id desc"
//	}
//	return p.Sort
//}
//func validHttpMethods(c *gin.Context) bool {
//	if c.Request.Method == "POST" || c.Request.Method == "GET" {
//		return true
//	} else {
//		emsg := "Not implement method:" + c.Request.Method
//		log.Logger.Error(emsg)
//		c.JSON(http.StatusOK, emsg)
//		return false
//	}
//}
