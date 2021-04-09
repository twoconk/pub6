package main

import (
	_ "gfast/boot"
	_ "gfast/router"

	"github.com/gogf/gf/frame/g"
)

// @title gfast API文档
// @version 1.0
// @description gfast 在线API文档
// @host localhost
// @BasePath /system
func main() {
	s := g.Server()
	// s.SetIndexFolder(false) //SetIndexFolder用来设置是否允许列出Server主目录的文件列表（默认为false）；
	// s.SetServerRoot("/home/www/")
	s.Run()
}
