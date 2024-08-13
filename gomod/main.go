package main

import (
	"encoding/xml"
	"fmt"
	"gomod/httpex"
	"gomod/mylib"
	"time"
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

// 定义一个结构体来包装 map 数据
type TimeMap struct {
	XMLName xml.Name `xml:"TimeMapxxx"`
	Entries []*Entry
}

// 定义一个结构体来表示 map 的每个条目
type Entry struct {
	Key   int32
	Value time.Time
}

func main2() {
	// 创建一个 map[int32]time.Time 类型的数据
	timex, _ := time.Parse("2006-01-02 15:04:05", "2024-06-01 00:00:00")
	timeMap := map[int32]time.Time{
		1: timex,
		2: timex.Add(time.Hour * 24),
	}

	// 将 map 数据转换为 TimeMap 结构体
	var entries []*Entry
	for k, v := range timeMap {
		entries = append(entries, &Entry{Key: k, Value: v})
	}
	timeMapStruct := TimeMap{Entries: entries}

	// 序列化为 XML
	xmlData, err := xml.MarshalIndent(timeMapStruct, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}

	// 打印 XML 数据
	fmt.Println(string(xmlData))
}

func main() {
	fmt.Println("hello world begin ! -------------")

	// mylib.ContainerTest()
	// mylib.TimeTest()
	// ostest.OsTest()
	// filerw.FileRW()

	// reflecttest.ReflectTest()
	mylib.TimeParse()

	httpex.UrlEncode()
	httpex.UrlDecode()

	testmap := make(map[int]int)
	value := testmap[1]
	if value == 0 {
		fmt.Println("is zero")
	}

}
