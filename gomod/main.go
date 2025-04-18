package main

import (
	_ "embed"
	"gomod/httpex"
	"log"
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
	go func() {
		httpex.WebSocketProxy()
	}()
	httpex.WebSocketServer()
	log.Println("Server ended")
}
