//Package 路由
package router

import (
	"github.com/gin-gonic/gin"
	. "farmer/autocs/apis"
	"net/http"
	"farmer/autocs/web"
	"farmer/autocs/admin"
	"farmer/autocs/middleware"

)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.Static("/upload", "./upload")
	router.LoadHTMLGlob("templates/**/*")	//定义模板路径

	router.GET("/", IndexApi)
	router.GET("/chat", ChatApi)

	router.GET("/chat_demo", func(c *gin.Context) {
		ChatDemoApi(c.Writer, c.Request)
	})

	router.POST("/user", AddPersonApi)

	router.GET("/users", GetPersonsApi)

	router.GET("/user/:id", GetPersonApi)

	//router.PUT("/person/:id", ModPersonApi)

	router.DELETE("/user/:id", DelPersonApi)

	router.GET("/info/:id",web.GetQaInfo)

	//api路由分组
	api := router.Group("/api")
	api.POST("/qa", CutWordsApi)
	api.GET("/qa/:id", GetQaInfoApi)


	//demo分组
	//demo := router.Group("demo")

	router.POST("/upload", UploadApi)


	// todo	登录处理和注销单独写在自定义函数里
	router.GET("/login", func(c *gin.Context) {
		//如果已经存在登录信息直接进入
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			if value != ""{
				c.Redirect(http.StatusFound, "/admin")
				c.Abort()
			}
		}
		c.HTML(http.StatusOK, "public/login.html", gin.H{"title": "Login Page", })
	})

	router.GET("/logout", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:"session_id",
			Value:"",
			Path:"/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.Redirect(http.StatusFound, "/login")
	})

	/**
	 desc: 登录验证接口
	 param:
	 	username string 用户名
	 	password string 密码
	 return:
	 	code	int		操作状态码	0 成功,1用户名错误, 2密码错误
	 	msg		string	提示信息
	 	url		string	跳转地址
	 */
	router.POST("/login", PublicLoginApi)


	/***** Admin分组路由 *********/
	adm := router.Group("/admin", middleware.AdminAuth())

	//管理后台
	adm.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{"title": "Main website","mid":"home"})
	})
	//	/admin/index采用301重定向到 /admin
	adm.GET("index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin")
	})

	//常见问题列表
	adm.GET("/faq", admin.QaList)
	//新增qa
	adm.GET("/faq/add", admin.QaAdd)
	adm.POST("/faq/add", admin.QaAddDo)
	//用户词典
	adm.GET("/dic", middleware.AdminAuth(), admin.GetAllWords)

	//问题详情
	router.GET("/faq/:id", admin.QaInfo)


	return router
}

