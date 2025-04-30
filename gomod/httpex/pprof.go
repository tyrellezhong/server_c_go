package httpex

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

func PprofServer() {
	// 启动 pprof 路由
	r := http.NewServeMux()
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))

	// 启动 HTTP 服务器
	addr := ":6060" // 你可以根据需要修改这个地址
	fmt.Printf("Starting pprof server at http://%s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}

func AutoProfServer() {
	go func() {
		log.Println("listening on port 6060")
		log.Fatal(http.ListenAndServe(":6060", nil))
	}()

	count := 0
	for {
		log.Printf("time pass %d minutes", count)
		time.Sleep(time.Second)
		count++
	}

}
