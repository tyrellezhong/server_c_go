package main

import (
	_ "embed"
	"log"
	"runtime"

	// _ "net/http/pprof" // 确保这里是匿名导入

	txttemplate "gomod/txt_template"
)

func main() {
	// 最多使用一个cpu
	runtime.GOMAXPROCS(1)

	// 设置日志格式，包含文件名和行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	txttemplate.ConditionTest()

	// httpex.PprofServer()
	// httpex.GinServer()
	// go httpex.HttpGetToFile("http://127.0.0.1:39001/debug/pprof/profile", "data_agent.profile.out")
	// for {
	// 	time.Sleep(5 * time.Second)
	// 	log.Println("sleep 5s")
	// }
	// filerw.FileRW()
	// flagex.PrintFlagArgs()
	// interfaceex.InterfaceTest()
	// zlog.NewCustomZLog()
	// httpex.UrlEncode()
	// flagex.PrintFlagArgs()
	queue := &[]int{1, 2, 3}
	count := 100
	for len(*queue) > 0 {
		log.Println((*queue)[0], "len(queue):", len(*queue))
		*queue = (*queue)[1:]
		*queue = append(*queue, count)
		count++
		if count > 110 {
			break
		}
	}
}
