package http

import (
	"github.com/gin-gonic/gin"
)

type QRequest struct {
	PageIndex int            `json:"PageIndex"`
	PageSize  int            `json:"PageSize"`
	Target    string         `json:"Target"` //目标对象
	Where     map[string]any `json:"Where"`
}
type ARequest struct {
	Target       string            `json:"Target"` //目标对象名称
	Fields       map[string]string `json:"Fields"`
	ObjectString string            `json:"ObjectString"`
}

type URequest struct {
	Target string         `json:"Target"` //目标对象名称
	Fields map[string]any `json:"Fields"`
	Where  map[string]any `json:"Where"`
}

type DRequest struct {
	Target string         `json:"Target"` //目标对象名称
	Where  map[string]any `json:"Where"`
}
type QResponse struct {
	PageIndex  int    `json:"PageIndex"`
	PageSize   int    `json:"PageSize"`
	TotalPages int    `json:"TotalPages"`
	Message    string `json:"message"`
	Code       int    `json:"code"`
	Data       any    `json:"Data"`
}
type HandleQueryTarget func(para *QRequest, c *gin.Context)
type HandleCreateTarget func(para *ARequest, c *gin.Context)
type HandleDeleteTarget func(para *DRequest, c *gin.Context)
type HandleUpdateTarget func(para *URequest, c *gin.Context)

type QueryTanslate[T any] interface {
	Translate(raw []T) []T
}

//const (
//	DefPageSize = 20
//)

//type Page[T any] struct {
//	PageSize   int    `json:"pageSize,omitempty" form:"pageSize"`
//	PageIndex  int    `json:"pageIndex,omitempty" form:"pageIndex"`
//	Sort       string `json:"sort,omitempty" form:"sort"`
//	TotalRows  int64  `json:"total_rows"`
//	TotalPages int    `json:"total_pages"`
//	Rows       []T    `json:"rows"`
//}

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
//func ValidHttpMethods(c *gin.Context) bool {
//	if c.Request.Method == "POST" || c.Request.Method == "GET" {
//		return true
//	} else {
//		emsg := "Not implement method:" + c.Request.Method
//		log.Logger.Error(emsg)
//		c.JSON(http.StatusOK, emsg)
//		return false
//	}
//}
