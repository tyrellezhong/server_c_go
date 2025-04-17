package main

import (
	_ "embed"
	"gomod/httpex"
	// txttemplate "gomod/txt_template"
)

func main() {

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
	httpex.HttpSvr()
}
