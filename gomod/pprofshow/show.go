package pprofshow

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// 采样表
const (
	TableProfile      = "profile"
	TableMemory       = "memory"
	TableBlock        = "block"
	TableMutex        = "mutex"
	TableGoroutine    = "goroutine"
	TableThreadcreate = "threadcreate"
	TableTrace        = "trace"
	TableProfHistory  = "profhistory"
)

type ProfShowInfo struct {
	ShowType   string
	Cmd        *exec.Cmd
	ListenAddr string
	StartTime  time.Time
	FileName   string
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

// CheckShutdown
func CheckShowShutdown() {
	ticker := time.NewTicker(time.Second * 60)
	if ticker != nil {
		go func() {
			for range ticker.C {
				log.Println("helloworld")
			}
		}()
	}
}

// CreatePprofServer
func CreateShowServer(info *ProfShowInfo) {
	// 定义要执行的命令和参数
	cmd := fmt.Sprintf("go tool pprof -http=%s %s", info.ListenAddr, info.FileName)
	info.Cmd = exec.Command("sh", "-c", cmd)

	msg, err := info.Cmd.Output()
	if err != nil {
		log.Printf("Error running command: %v", err)
		return
	}
	log.Printf("Command output: %s", msg)
}
