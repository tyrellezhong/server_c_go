package timeex

import (
	"fmt"
	"time"
)

func WaitChannel(conn <-chan string) {
	timer := time.NewTimer(10 * time.Second)
	select {
	case <-conn:
		fmt.Println("conn")
	case <-timer.C:
		fmt.Println("timeout")
	}
	<-time.After(1 * time.Second)
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("after func")
	})
}
