package main

import (
	_ "embed"
	"fmt"
	"gomod/filerw"
	"gomod/httpex"
)

func main() {
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

	// days := []gogenerate.Weekday{gogenerate.Sunday, gogenerate.Monday, gogenerate.Tuesday}
	// for _, day := range days {
	// 	fmt.Println(day.String())
	// }
	// 文件读写测试
	filerw.FileRW()
}
