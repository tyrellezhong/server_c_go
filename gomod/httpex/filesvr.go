package httpex

import (
	"fmt"
	"net/http"
)

func FileServer() {

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		// 使用 ServeFile 发送文件
		http.ServeFile(w, r, "path/to/your/file.txt")
	})

	// 指定静态文件目录
	fs := http.FileServer(http.Dir("../web_source/webtest"))

	// 将文件服务器挂载到 "/static" 路径
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 启动 HTTP 服务器
	fmt.Printf("Starting pprof server at http://%s\n", ":8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
