package main

import (
	_ "embed"
	"gomod/httpex"
	"log"

	// _ "net/http/pprof" // 确保这里是匿名导入
	"time"
	// txttemplate "gomod/txt_template"
)

func main() {

	// 设置日志格式，包含文件名和行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// txttemplate.ConditionTest()
	errChan := make(chan error, 1)
	go func() {
		errChan <- nil
	}()
	err := <-errChan
	if err != nil {
		println("err:", err)
		return
	} else {
		println("err is nil")
	}
	println("err:", err)
	// go func() {
	// 	log.Println("listening on port 6060")
	// 	log.Fatal(http.ListenAndServe(":6060", nil))
	// }()
	httpex.ProffServer()

	count := 0
	for {
		log.Printf("time pass %d minutes", count)
		time.Sleep(time.Second)
		count++
	}

	log.Println("Server ended")
}
