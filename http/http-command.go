package http

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"wilder.cn/gogo/log"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
type DefaultResponse[T interface{}] struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    T      `json:"data"`
}
type QueryRequest struct {
	PageIndex int            `json:"PageIndex"`
	PageSize  int            `json:"PageSize"`
	Code      int            `json:"Code"` //操作码
	Where     map[string]any `json:"Where"`
	SqlItems  []interface{}  `json:"SqlItems"`
}
type DeleteRequest struct {
	Code  int            `json:"Code"` //操作码
	Where map[string]any `json:"Where"`
}
type GetRequest struct {
	Code        int    `json:"Code"` //操作码
	FilterName  string `json:"FilterName" binding:"required"`
	FilterValue string `json:"FilterValue" binding:"required"`
}
type UpdateRequest struct {
	Code   int            `json:"Code"` //操作码
	Fields map[string]any `json:"Fields"`
	Where  map[string]any `json:"Where"`
}

func (r UpdateRequest) LookField(key string) (string, bool) {
	v, ok := r.Fields[key]
	if ok {
		return fmt.Sprintf("%v", v), true
	} else {
		return "", false
	}
}

func (r UpdateRequest) TryField(key string) (string, error) {
	v, ok := r.Fields[key]
	if ok {
		return fmt.Sprintf("%v", v), nil
	} else {
		return "", NewMVCError(ErrorParaNotExist, key+" not found in request.")
	}
}

type QueryResponse[T any] struct {
	PageIndex  int    `json:"PageIndex"`
	PageSize   int    `json:"PageSize"`
	TotalPages int    `json:"TotalPages"`
	TotalRows  int64  `json:"TotalRows"`
	Message    string `json:"message"`
	Code       int    `json:"code"`
	Data       []*T   `json:"Data"`
}

type WError interface {
	error
	Code() int
}

func RequestToPage[T any](req QueryRequest) *Page[T] {
	return &Page[T]{
		PageSize:  req.PageSize,
		PageIndex: req.PageIndex,
	}
}
func PageToResponse[T any](res *Page[T]) *QueryResponse[T] {
	return &QueryResponse[T]{
		PageSize:   res.PageSize,
		PageIndex:  res.PageIndex,
		TotalPages: res.TotalPages,
		TotalRows:  res.TotalRows,
		Data:       res.Rows,
	}
}

type AuthService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

func apiResponse(message string, code int, data interface{}) Response {

	jsonResponse := Response{
		Message: message,
		Code:    code,
		Data:    data,
	}
	return jsonResponse
}
func SuccessResponse(data interface{}) Response {
	rep := apiResponse("success", RSuccess, data)
	return rep
}
func FailedResponse(err string, data interface{}) Response {
	rep := apiResponse(err, RFailed, data)
	return rep
}
func FailedResponseCode(ec int, err string, data interface{}) Response {
	rep := apiResponse(err, ec, data)
	return rep
}
func QueryPage[T any](para *QueryRequest, c *gin.Context) {
	log.Logger.InfoF("Received query:%s", c.Request.RequestURI)
	con := para.Where
	log.Logger.InfoF("query map:%v", *para)

	pg := NewPage[T]()
	pg.PageSize = para.PageSize
	pg.PageIndex = para.PageIndex
	res, err := SelectPage(con, pg)
	if err != nil {
		c.JSON(http.StatusOK, FailedResponse("query failed", err))
	} else {
		qres := &QueryResponse[T]{}
		qres.PageSize = res.PageSize
		qres.PageIndex = res.PageIndex
		qres.TotalPages = res.TotalPages
		qres.Data = res.Rows
		c.JSON(http.StatusOK, SuccessResponse(qres))
	}
}
func HandlErr(c *gin.Context, err error) {
	errs, ok := Format(err)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		c.JSON(http.StatusOK, FailedResponse("Server internal error.", err.Error()))
		return
	}
	// validator.ValidationErrors类型错误则进行翻译
	c.JSON(http.StatusOK, FailedResponseCode(CommParaFormat, "valid failed", errs))
	return
}
