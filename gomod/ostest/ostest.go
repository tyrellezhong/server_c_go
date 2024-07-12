package ostest

import (
	"fmt"
	"os"
)

func OsTest() {
	hostName, _ := os.Hostname()
	fmt.Println("hostname:", hostName)

}
