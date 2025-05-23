package httpex

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
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

// HttpGetToFile http请求url地址（入参），并将返回结果写入文件outputfile（入参）
func HttpGetToFile(url string, outputfile string) {
	// 创建一个HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 创建一个HTTP客户端
	client := &http.Client{}

	// 发送请求并获取响应
	fmt.Printf("Sending request... url=%s", url)
	resp, err := client.Do(req)
	fmt.Printf("Response received resp=%v", resp.Status)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容并写入文件
	file, err := os.Create(outputfile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error copying response body to file:", err)
		return
	}
}

// CreatePprofServer
