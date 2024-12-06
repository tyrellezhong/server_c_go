package httpex

import (
	"fmt"
	"net/http"
	"net/url"
)

func UrlEncode() {
	// 创建一个 url.Values 对象
	values := url.Values{}

	tmpName := "=John Doe& ?"
	// 添加查询参数
	values.Add("name", tmpName)
	values.Add("age", "30")
	values.Add("hobby", "reading")
	values.Add("hobby", "travelling")

	nameEncode := url.QueryEscape(tmpName)
	fmt.Println("nameEncode:", nameEncode)
	nameUnencode, _ := url.QueryUnescape(nameEncode)
	fmt.Println("nameUnencode", nameUnencode)

	// 编码为查询字符串
	queryString := values.Encode()
	fmt.Println("Query string:", queryString)
	// 输出: Query string: age=30&hobby=reading&hobby=travelling&name=John+Doe
	queryValue, _ := url.ParseQuery(queryString)

	name := queryValue.Get("name")
	fmt.Printf("name:%s\n", name)

	// 构建完整的 URL
	baseURL := "https://example.com/search"
	fullURL := baseURL + "?" + queryString
	fmt.Println("Full URL:", fullURL)
	// 输出: Full URL: https://example.com/search?age=30&hobby=reading&hobby=travelling&name=John+Doe
}

func UrlDecode() {
	// 解析一个 URL
	urlstr1 := "https://example.com/search?name=John+Doe&age=30&hobby=reading&hobby=travelling&special=key+with+spaces+%26+special%3Dcharacters%3F"

	u, err := url.Parse(urlstr1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 获取查询参数
	values := u.Query()

	// 访问查询参数
	name := values.Get("name")
	age := values.Get("age")
	hobbies := values["hobby"]
	special := values.Get("special")

	fmt.Println("Name:", name)       // 输出: Name: John Doe
	fmt.Println("Age:", age)         // 输出: Age: 30
	fmt.Println("Hobbies:", hobbies) // 输出: Hobbies: [reading travelling]
	fmt.Println("special:", special)

}

func HandlerRedirect(w http.ResponseWriter, r *http.Request) {
	// 设置响应头的 Content-Type 为 text/html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 构建 HTML 内容，包含自动跳转的 meta 标签
	htmlContent := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="refresh" content="5;url=https://www.google.com">
        <title>自动跳转示例</title>
    </head>
    <body>
        <h1>欢迎来到我的网站!</h1>
        <p>您将在 5 秒后自动跳转到 Google。</p>
        <p>如果没有自动跳转，请点击 <a href="https://www.google.com">这里</a>.</p>
    </body>
    </html>
    `

	// 写入响应
	fmt.Fprint(w, htmlContent)
}

func HandlerRedirect2(w http.ResponseWriter, r *http.Request) {
	// 使用 HTTP 302 重定向到目标 URL
	http.Redirect(w, r, "https://www.google.com", http.StatusFound)
}

func HttpSvr() {
	// 设置路由和处理函数
	http.HandleFunc("/", HandlerRedirect)
	http.HandleFunc("/redirect", HandlerRedirect2)

	// 启动 HTTP 服务器
	fmt.Println("服务器正在运行，访问 http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}
