package mylib

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Sum(a, b int) int {
	return a + b
}

func ContainerTest() {
	SliceTest()
	StringTest()
	MapTest()
	ListTest()
}

func SliceTest() {
	fmt.Println("SliceTest ------------------------------")
	var slice = make([]int, 5, 50)
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3, 9, 8, 7)
	slice = slice[5:]
	sort.Ints(slice)

	for i, val := range slice {
		fmt.Println(i, val)
	}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(i, slice[i])
	// }
}

func StringTest() {
	fmt.Println("StringTest ------------------------------")

	str := "Hello, World!"
	str1 := "Hello"
	str2 := "World"

	// 使用 + 操作符连接字符串
	result := str1 + ", " + str2 + "!"
	fmt.Println(result)

	// 使用 strings.Join 连接字符串
	parts1 := []string{"Hello", "World"}
	result = strings.Join(parts1, ", ")
	fmt.Println(result)

	// 查找子字符串
	index := strings.Index(str, "World")
	fmt.Println(index) // 输出: 7

	// 检查子字符串是否存在
	contains := strings.Contains(str, "World")
	fmt.Println(contains) // 输出: true

	// 替换子字符串
	newStr := strings.Replace(str, "World", "Go", 1)
	fmt.Println(newStr) // 输出: Hello, Go!

	str3 := "a,b,c,d,e"

	// 分割字符串
	parts2 := strings.Split(str3, ",")
	fmt.Println(parts2) // 输出: [a b c d e]

	str4 := "Hello,世界!"

	// 使用 for range 遍历字符串
	for _, char := range str4 {
		fmt.Printf("%c\n", char)
	}
	// 获取字符串长度
	length := len(str4)
	fmt.Println(length) // 输出: 13 (字节长度)

	// 获取字符长度
	runeLength := len([]rune(str4))
	fmt.Println(runeLength) // 输出: 9 (字符长度)

	// 整数转换为字符串
	intVal := 123
	strVal := strconv.Itoa(intVal)
	fmt.Println(strVal) // 输出: 123

	// 字符串转换为整数
	str5 := "456"
	intVal, err := strconv.Atoi(str5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(intVal) // 输出: 456
	}

	str6 := "Hello,World!"

	// 将字符串转换为 []rune
	runes := []rune(str6)

	// 修改字符
	runes[7] = 'G'
	runes[8] = 'o'

	// 将 []rune 转换回字符串
	newStr2 := string(runes)
	fmt.Println(newStr2) // 输出: Hello, Go!

}

func MapTest() {
	fmt.Println("MapTest ------------------------------")
	var dic = make(map[int]string)
	dic[1] = "a"
	dic[2] = "b"
	delete(dic, 1)
	if val, ok := dic[2]; ok {
		fmt.Println(val)
	}
}

func ListTest() {
	fmt.Println("ListTest ------------------------------")
	li := list.New()
	li.PushBack("123")
	e := li.PushBack("456")
	li.InsertBefore("---", e)

	for n := li.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}

}

func TimeTest() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp := now.Unix()            //时间戳
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)

	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))

	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2024-07-17 5:00:00"
	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj2)

	// 加载不同的时区
	locNY, _ := time.LoadLocation("America/New_York")
	locLA, _ := time.LoadLocation("America/Los_Angeles")

	// 使用不同的时区解析时间
	tNY, _ := time.ParseInLocation(layout, timeStr, locNY)
	tLA, _ := time.ParseInLocation(layout, timeStr, locLA)
	tCN, _ := time.ParseInLocation(layout, timeStr, time.Local)
	utc, _ := time.Parse(layout, timeStr)
	ut2, _ := time.ParseInLocation(layout, timeStr, time.UTC)

	fmt.Println("utc time ", utc)
	fmt.Println("utc2 time ", ut2)
	fmt.Println("Parsed time (New York):", tNY)
	fmt.Println("Parsed time (Los Angeles):", tLA)
	fmt.Println("Parsed time (Cn):", tCN)

	// 比较两个时间
	fmt.Println("Is New York time before Los Angeles time?", tNY.Before(tLA))
	fmt.Println("Is New York time before CN time?", tNY.Before(tCN))
	fmt.Println("Is New York time before utc time?", tNY.Before(utc))
	fmt.Println("Is cn time before utc time?", tCN.Before(utc))
	futureTime := tNY.Add(time.Hour)
	fmt.Println("Is now time before future time? ", tNY.Before(futureTime), tNY, futureTime)
	fmt.Println("Is cn time after new york time? ", now.After(tNY), now, tNY)

}
