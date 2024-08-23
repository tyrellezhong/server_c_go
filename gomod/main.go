package main

import (
	_ "embed"
	"fmt"
	"gomod/gogenerate"
	"gomod/httpex"
	"os"
	"text/template"
)

//go:embed service.gotpl
var serviceTemplate string

func main() {
	// 解析模板
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// 定义模板数据
	data := struct {
		PackageName string
		ServiceName string
	}{
		PackageName: "example",
		ServiceName: "ExampleService",
	}

	// 执行模板并输出到标准输出
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func main2() {
	fmt.Println("hello world begin ! -------------")
	// riskInfo := &PayRickControlInfo{
	// 	AdultStatus:         -1,
	// 	RegisterRegion:      "\"KR\"",
	// 	PlayerAge:           13,
	// 	LastModifyTimestamp: 4,
	// }
	// logger, _ := zap.NewProduction()
	// defer logger.Sync()
	// logger.Info("pay risk control check", zap.Any("riskctl info", riskInfo))

	// mylib.ContainerTest()
	// mylib.TimeTest()
	// ostest.OsTest()
	// filerw.FileRW()

	// // reflecttest.ReflectTest()
	// mylib.TimeParse()

	httpex.UrlEncode()
	httpex.UrlDecode()

	// testmap := make(map[int]int)
	// value := testmap[1]
	// if value == 0 {
	// 	fmt.Println("is zero")
	// }

	days := []gogenerate.Weekday{gogenerate.Sunday, gogenerate.Monday, gogenerate.Tuesday}
	for _, day := range days {
		fmt.Println(day.String())
	}
}
