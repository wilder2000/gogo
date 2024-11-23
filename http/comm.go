package http

import (
	"github.com/gin-gonic/gin"
	hp "net/http"
	"wilder.cn/gogo/log"
)

const (
	RSuccess     = 0
	RFailed      = 1 //未知原因
	RNeedLogin   = 2 //需要登录
	RServerError = 3 //服务器内部错误

	LoginParaFormat = 21 //参数不正确
	PwdWrong        = 22 //密码不正确
	UserNotFound    = 23 //用户不存在
	DataExistFound  = 24 //数据已经存在
	CommParaFormat  = 25 //参数不正确

	CmdImagePrefixPage = 1100 //文档的链接前缀

	CmdUserGroups            = 1200 //查询用户加入的编组
	CmdQueryGroups           = 1201 //查询用户编组
	CmdQueryUnionGroups      = 1202 //查询用户加入的编组，并上没有加入的编组
	CmdQueryUnionUsers       = 1203 //查询编组下面的用户，并上没有加入的用户
	CmdQueryUnionRoles       = 1204 //查询编组加入的角色，并上没有加入的角色
	CmdQueryUnionOperators   = 1205 //查询角色拥有的操作权限，并上没有的操作权限
	CmdQueryUnionDepartments = 1206 //查询部门拥有的用户，并上没有加入的用户

	CmdQueryDocBaseDepartment = 1207 //查询文件柜信息，关联用户名称和部门名称
	CmdQueryDocBaseDocs       = 1208 //查询文件柜所有文档，关联用户名称,邮箱
	CmdQueryTheDoc            = 1209 //查询指定的文档，关联用户名称,邮箱
	CmdQueryTheDocTags        = 1210 //查询指定的文档的所有标签
	CmdQueryMyRecentUpDoc     = 1211 //查询指定用户最近上传的文档
	CmdQueryDBRecentUpDoc     = 1212 //查询指定文件柜最近上传的文档
	CmdQueryDBAllTags         = 1213 //查询指定文件所有用到的标签
	CmdQueryDBTagsDocs        = 1214 //查询文件柜指定标签所有文档，关联用户名称,邮箱
	CmdQueryDBList            = 1215 //查询文件柜列表，关联用户名称,邮箱
	CmdDeleteDB               = 1216 //删除指定文件柜
	CmdQueryDocListQuery      = 1217 //文档高级查询，名称模糊匹配，同时可以指定部门、文件柜、标签、时间范围
	CmdDeleteDoc              = 1218 //删除指定文件
	CmdQueryDocDepPublic      = 1219 //查询文件柜关联的部门和非关联的部门
	CmdQueryDocDepPrivate     = 1220 //查询文件柜关联的部门

	CmdUpdateUserName = 1207 //修改用户的别名
	CmdUpdateUserSex  = 1208 //修改用户的性别
	CmdGetUserBasic   = 1209 //获取用户的基础信息

	//UserStateAuto     = 2 //自动注册
	UserStateEnable   = 1 //启用
	UserStateDisable  = 0 //禁用
	ErrorParaNotExist = 31
)

type MVCError struct {
	errcode int
	message string
}

func NewMVCError(ec int, msg string) *MVCError {
	err := &MVCError{
		errcode: ec,
		message: msg,
	}
	return err
}
func (M MVCError) Error() string {
	return M.message
}

func (M MVCError) Code() int {
	return M.errcode
}

type HTTPController[T any] interface {
	Execute(para *T, c *gin.Context)
	Prepare(c *gin.Context)
	UrlPath() string
}
type AbstractController[T any] struct {
	HTTPController[T]
}

func (b *AbstractController[T]) Execute(para *T, c *gin.Context) {
	log.Logger.InfoF("HTTPController %+v", b.HTTPController)
	log.Logger.InfoF("HTTPController Execute %+v", b.HTTPController.Execute)
	b.HTTPController.Execute(para, c)
}

func (b *AbstractController[T]) Prepare(c *gin.Context) {
	var paraModel T

	if err := c.ShouldBind(&paraModel); err == nil {
		//自己实现一个统一的controler,再分发，再简化controler的实现
		if c.Request.Method == "POST" || c.Request.Method == "GET" {
			b.Execute(&paraModel, c)
			return
		} else {
			emsg := "Not implement method:" + c.Request.Method
			log.Logger.Error(emsg)
			c.JSON(hp.StatusOK, emsg)
		}

	} else {
		errs, ok := Format(err)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(hp.StatusOK, FailedResponse("Server internal error.", err.Error()))
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(hp.StatusOK, FailedResponseCode(CommParaFormat, "valid failed", errs))
		return
	}

}
