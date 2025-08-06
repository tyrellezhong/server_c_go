package flagex

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	// 支持输入
	// -name=myname
	// -name myname
	// --name=myname
	// --name myname
	ArgBase64 = flag.String("base64", "", "base64编码的字符串")
	ArgOct    = flag.String("oct", "", "八进制显示的字符串")
	// 除了以上四种，还支持 -bool (解析为true)
	ArgBool = flag.Bool("bool", false, "input bool")
)

// 判断字符是否为八进制数字
func isOctalDigit(c byte) bool {
	return c >= '0' && c <= '7'
}

// 将八进制转义序列转换为 UTF-8 字符串
func octalToUTF8(octStr string) string {
	var utf8Str strings.Builder

	for i := 0; i < len(octStr); {
		if octStr[i] == '\\' && i+3 < len(octStr) && isOctalDigit(octStr[i+1]) && isOctalDigit(octStr[i+2]) && isOctalDigit(octStr[i+3]) {
			// 提取八进制数字
			octNum := octStr[i+1 : i+4]
			// 将八进制数字转换为十进制
			decNum, err := strconv.ParseInt(octNum, 8, 32)
			if err != nil {
				fmt.Println("Error parsing octal number:", err)
				return ""
			}
			// 将十进制数字转换为字符并追加到结果字符串
			utf8Str.WriteByte(byte(decNum))
			// 跳过已处理的字符
			i += 4
		} else {
			// 处理非八进制转义字符
			utf8Str.WriteByte(octStr[i])
			i++
		}
	}

	return utf8Str.String()
}

func PrintFlagArgs() {
	flag.Parse()
	var utf8Str string
	encodedStr := *ArgBase64
	var parsed = false
	if encodedStr == "" {
		encodedStr = os.Args[1]
	}
	if len(*ArgBase64) > 0 {
		decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
		if err != nil {
			log.Fatalf("Failed to decode Base64 string: %v", err)
		}
		utf8Str = string(decodedBytes)
		fmt.Println("base64 str to utf8: ", utf8Str)
		parsed = true
	}
	if len(*ArgOct) > 0 {
		utf8Str = octalToUTF8(*ArgOct)
		fmt.Println("oct str to utf8: ", utf8Str)
		parsed = true
	}

	if len(os.Args[1]) > 0 && !parsed {
		decodedBytes, err := base64.StdEncoding.DecodeString(os.Args[1])
		if err == nil {
			utf8Str = string(decodedBytes)
			fmt.Println("base64 str to utf8: ", utf8Str)
		} else {
			utf8Str = octalToUTF8(os.Args[1])
			fmt.Println("oct str to utf8: ", utf8Str)
		}
	}
}
