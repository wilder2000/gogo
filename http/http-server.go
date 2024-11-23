package http

//
//var (
//	mappings map[string]gin.HandlerFunc
//)
//
//func InitController(e *gin.Engine) {
//	uh := NewUserHandler(UserProxy)
//	e.POST("/login", uh.Login)
//	e.POST("/reg", uh.RegisterUser)
//	proUrlGrp := e.Group("/v1", PreProcess)
//	for path, hl := range mappings {
//		proUrlGrp.POST(path, hl)
//		log.Logger.InfoF("Mapping:%s", path, reflect.TypeOf(hl).Name())
//	}
//	proUrlGrp.POST(mif.REQCreate, mif.HandleCreate)
//	proUrlGrp.POST(mif.REQQuery, mif.HandleQuery)
//	proUrlGrp.POST(mif.REQDelete, mif.HandleDelete)
//	proUrlGrp.POST(mif.REQUpdate, mif.HandleUpdate)
//}
//
//func Start(port string, readout time.Duration, wout time.Duration, actions ...HttpController) {
//	router := gin.Default()
//	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
//		// your custom format
//		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
//			param.ClientIP,
//			param.TimeStamp.Format(time.RFC1123),
//			param.Method,
//			param.Path,
//			param.Request.Proto,
//			param.StatusCode,
//			param.Latency,
//			param.Request.UserAgent(),
//			param.ErrorMessage,
//		)
//	}))
//	router.Use(gin.Recovery())
//	InitController(router)
//	for _, mapping := range actions {
//		router.POST(mapping.Path, mapping.Action)
//		log.Logger.InfoF(" POST Mapping:%s", mapping.Path)
//	}
//	srv := &http.Server{
//		Addr:           port,
//		Handler:        router,
//		ReadTimeout:    readout * time.Second,
//		WriteTimeout:   wout * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//	go func() {
//		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
//			log.Logger.InfoF("http server error: %s\n", err)
//		}
//	}()
//	log.Logger.InfoF("GOGO Http Server started success.")
//	// Wait for interrupt signal to gracefully shutdown the server with
//	// a timeout of 5 seconds.
//	quit := make(chan os.Signal)
//	// kill (no param) default send syscall.SIGTERM
//	// kill -2 is syscall.SIGINT
//	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//	log.Logger.Info("Shutting down server...")
//	// The context is used to inform the server it has 5 seconds to finish
//	// the request it is currently handling
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Logger.ErrorF("Server forced to shutdown:", err)
//	}
//
//	log.Logger.Info("Server exiting")
//}
//
//type HttpController struct {
//	Path   string
//	Action gin.HandlerFunc
//}
