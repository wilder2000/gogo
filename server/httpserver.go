package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	gh "net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"
	"wilder.cn/gogo/config"
	"wilder.cn/gogo/http"
	"wilder.cn/gogo/log"
	"wilder.cn/gogo/sys/mif"
	"wilder.cn/gogo/sys/sm"
)

func init() {
	fmt.Printf("%s\n", config.Logo)
	fmt.Printf("%s\n", config.LogoTitle)
	RegMapping[http.ChangePWD](new(sm.PwdController))
	RegMapping[http.CheckPWD](new(sm.CheckPwdController))
	RegMapping[http.QueryRequest](new(sm.UserQueryController))
	RegMapping[http.QueryRequest](new(sm.RoleQueryController))
	RegMapping[sm.AddRoleRequest](new(sm.RoleAddController))
	RegMapping[http.QueryRequest](new(sm.UserGroupsController))
	RegMapping[http.QueryRequest](new(sm.OperatorController))
	RegMapping[http.QueryRequest](new(sm.DepartmentController))
	RegMapping[http.UpdateRequest](new(sm.UserMgrController))
	RegMapping[http.GetRequest](new(sm.UserProfileController))

}

var (
	mappings       = make(map[string]gin.HandlerFunc)
	noAuthMappings = make(map[string]gin.HandlerFunc)
)

func InitController(e *gin.Engine) *gin.RouterGroup {
	uh := http.NewUserHandler(http.UserProxy)
	e.POST("/api/emllogin", uh.EmailLogin)
	if config.AConfig.Security.Registration {
		e.POST("/api/reguser", uh.RegisterUser)
	}

	e.POST("/api/reqmcode", uh.RequestMobileCode)
	e.POST("/api/token_valid", uh.FileAccessValid)
	//e.POST("/autoreguser", uh.AutoRegUser)
	e.POST("/api/updmobile", uh.UpdateMobile)

	e.POST("/api/moblogin", uh.MobileLogin)
	e.POST("/api/newreglogin", uh.UIDLoginRegist)
	e.POST("/api/loginexist", uh.UIDLoginWithExist)
	for path, hl := range noAuthMappings {
		e.POST(path, hl)
		log.Logger.InfoF("NO Auth Mapping:%s Function:%s", path, reflect.TypeOf(hl).Name())
	}

	proUrlGrp := e.Group("/api/v1", http.PreProcess)
	for path, hl := range mappings {
		proUrlGrp.POST(path, hl)
		log.Logger.InfoF("Mapping:%s Function:%s", path, reflect.TypeOf(hl).Name())
	}
	proUrlGrp.POST("/avatorup", uh.UploadAvatar)
	proUrlGrp.POST("/requestuser", uh.RequestUser)
	proUrlGrp.POST("/delaccount", uh.DeleteAccount)
	proUrlGrp.POST("/modalias", uh.UpdateAliasName)
	proUrlGrp.POST("/reperror", uh.ReportErrors)
	proUrlGrp.POST(mif.REQCreate, mif.HandleCreate)
	proUrlGrp.POST(mif.REQQuery, mif.HandleQuery)
	proUrlGrp.POST(mif.REQDelete, mif.HandleDelete)
	proUrlGrp.POST(mif.REQUpdate, mif.HandleUpdate)
	return proUrlGrp
}

// Start http server start func
func Start(address string, readout time.Duration, wout time.Duration, actions []HttpController, noauthActions []HttpController) {

	router := gin.Default()
	staticDir := config.AConfig.StaticDir
	//web := config.AConfig.Web
	if web := config.AConfig.Web; web != nil {
		for k, v := range web {
			fmt.Printf("web static :%s -> %s\n", k, v)
			router.Static(k, v)
		}
	}
	if len(strings.TrimSpace(staticDir.AbsoluteFileDir)) > 0 {
		router.Static(staticDir.RelativePath, staticDir.AbsoluteFileDir)
		log.Logger.InfoF("mapping www url:%s file dir:%s", staticDir.RelativePath, staticDir.AbsoluteFileDir)
	} else {
		log.Logger.InfoF("no mapping www path config.")
	}
	if len(strings.TrimSpace(staticDir.AbsoluteFileDir2)) > 0 {
		router.Static(staticDir.RelativePath2, staticDir.AbsoluteFileDir2)
		log.Logger.InfoF("mapping www url:%s file dir:%s", staticDir.RelativePath2, staticDir.AbsoluteFileDir2)
	} else {
		log.Logger.InfoF("no mapping www path config.")
	}
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	if config.AConfig.AccessControlAllowOrigin {
		log.Logger.InfoF("AccessControlAllowOrigin true")
		log.Logger.InfoF("AllowHost: ", config.AConfig.AccessControlAllowHost)
		log.Logger.InfoF("AllowMethods: ", config.AConfig.AccessControlAllowMethods)
		log.Logger.InfoF("AllowHeaders: ", config.AConfig.AccessControlAllowHeaders)
		router.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", config.AConfig.AccessControlAllowHost) // 允许任何源
			c.Writer.Header().Set("Access-Control-Allow-Methods", config.AConfig.AccessControlAllowMethods)
			c.Writer.Header().Set("Access-Control-Allow-Headers", config.AConfig.AccessControlAllowHeaders)
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return // 直接返回204状态码
			}
			c.Next() // 继续执行其他的中间件链
		})
	} else {
		log.Logger.InfoF("AccessControlAllowOrigin false")
	}

	for _, mapping := range noauthActions {
		noAuthMappings[mapping.Path] = mapping.Action
	}
	rr := InitController(router)
	for _, mapping := range actions {
		log.Logger.InfoF(" POST Mapping:%s", mapping.Path)
		rr.POST(mapping.Path, mapping.Action)
	}

	srv := &gh.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    readout * time.Second,
		WriteTimeout:   wout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, gh.ErrServerClosed) {
			log.Logger.InfoF("http server error: %s\n", err)
		} else {
			log.Logger.InfoF("GOGO Http Server started success. Binding :%s", address)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutting down server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.ErrorF("Server forced to shutdown:", err)
	}

	log.Logger.Info("Server exiting")
}

type HttpController struct {
	Path   string
	Action gin.HandlerFunc
}

func RegMapping[M any](c http.HTTPController[M]) {
	ctrl := newController(c)
	mappings[c.UrlPath()] = ctrl.Prepare
}

func newController[M any](c http.HTTPController[M]) *http.AbstractController[M] {
	ty := reflect.ValueOf(c)
	fi := ty.Elem().FieldByName("AbstractController")
	if fi.Type().ConvertibleTo(reflect.TypeOf(http.AbstractController[M]{})) {
		cc := fi.Interface().(http.AbstractController[M])
		cc.HTTPController = c
		return &cc
	}
	return nil
}
