package main

import (
	"fmt"
	"gomod/httpex"
)

// "github.com/gin-gonic/gin"

func getValues() (int, int) {
	return 10, 20
}

// func main() {
// 	x := 5 // 已经声明的变量

// 	// 短变量声明，同时赋值给已经声明的变量 x 和新变量 y
// 	x, y := getValues()

// 	fmt.Println("x:", x) // 输出: x: 10
// 	fmt.Println("y:", y) // 输出: y: 20
// }

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

}
