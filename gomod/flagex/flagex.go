package flagex

import (
	"flag"
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
}
