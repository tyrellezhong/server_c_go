package flagex

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// 支持输入
	// -name=myname
	// -name myname
	// --name=myname
	// --name myname
	ArgName = flag.String("name", "", "input name")
	// 除了以上四种，还支持 -bool (解析为true)
	ArgBool = flag.Bool("bool", false, "input bool")
)

func PrintFlagArgs() {
	flag.Parse()
	println("name:", *ArgName)
	println("bool:", *ArgBool)

	// 这是你提供的 Base64 编码的字符串
	encodedStr := *ArgName
	if encodedStr == "" {
		encodedStr = os.Args[1]
	}
	println("encodedStr:", encodedStr)

	// Base64 解码
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		log.Fatalf("Failed to decode Base64 string: %v", err)
	}

	// 将解码后的字节转换为 UTF-8 字符串
	utf8String := string(decodedBytes)

	// 打印还原的 UTF-8 字符串
	fmt.Println("Decoded UTF-8 string:", utf8String)
}
