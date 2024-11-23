package sm

//type HTTPController[T any] interface {
//	Execute(para *T, c *gin.Context)
//	Prepare(c *gin.Context)
//	UrlPath() string
//}

//type AbstractController[T any] struct {
//	HTTPController[T]
//}

//func (b *AbstractController[T]) Execute(para *T, c *gin.Context) {
//	b.HTTPController.Execute(para, c)
//}

//func (b *AbstractController[T]) Prepare(c *gin.Context) {
//	var paraModel T
//
//	if err := c.ShouldBind(&paraModel); err == nil {
//		//自己实现一个统一的controler,再分发，再简化controler的实现
//		if c.Request.Method == "POST" || c.Request.Method == "GET" {
//			b.Execute(&paraModel, c)
//			return
//		} else {
//			emsg := "Not implement method:" + c.Request.Method
//			log.Logger.Error(emsg)
//			c.JSON(hp.StatusOK, emsg)
//		}
//
//	} else {
//		errs, ok := http.Format(err)
//		if !ok {
//			// 非validator.ValidationErrors类型错误直接返回
//			c.JSON(hp.StatusOK, http.FailedResponse("Server internal error.", err.Error()))
//			return
//		}
//		// validator.ValidationErrors类型错误则进行翻译
//		c.JSON(hp.StatusOK, http.FailedResponseCode(http.CommParaFormat, "valid failed", errs))
//		return
//	}
//
//}
