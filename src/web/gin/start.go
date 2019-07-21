package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start() {
	//初始化一个引擎
	app:=gin.Default()
	//注册一个路由和处理函数
	app.Any("/", webroot)
	//绑定端口并启动运用
	app.Run(":12345")
	app.Run()
}

func webroot(context *gin.Context) {
	context.String(http.StatusOK, "This is /.")
}
