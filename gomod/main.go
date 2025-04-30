package main

import (
	_ "embed"
	"gomod/httpex"
	"log"
	// _ "net/http/pprof" // 确保这里是匿名导入
	// txttemplate "gomod/txt_template"
)

func main() {

	// 设置日志格式，包含文件名和行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// txttemplate.ConditionTest()

	// httpex.PprofServer()
	// httpex.GinServer()
	httpex.FileServer()
}
