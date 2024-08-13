package httpex

import (
	"fmt"
	"net/url"
)

func UrlEncode() {
	// 创建一个 url.Values 对象
	values := url.Values{}

	// 添加查询参数
	values.Add("name", "John Doe")
	values.Add("age", "30")
	values.Add("hobby", "reading")
	values.Add("hobby", "travelling")

	// 编码为查询字符串
	queryString := values.Encode()
	fmt.Println("Query string:", queryString)
	// 输出: Query string: age=30&hobby=reading&hobby=travelling&name=John+Doe
	queryValue, err := url.ParseQuery(queryString)
	name := queryValue.Get("name")
	fmt.Printf("name:%s\n%v", name, err)

	// 构建完整的 URL
	baseURL := "https://example.com/search"
	fullURL := baseURL + "?" + queryString
	fmt.Println("Full URL:", fullURL)
	// 输出: Full URL: https://example.com/search?age=30&hobby=reading&hobby=travelling&name=John+Doe
}

func UrlDecode() {
	// 解析一个 URL
	u, err := url.Parse("https://example.com/search?name=John+Doe&age=30&hobby=reading&hobby=travelling")
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

	fmt.Println("Name:", name)       // 输出: Name: John Doe
	fmt.Println("Age:", age)         // 输出: Age: 30
	fmt.Println("Hobbies:", hobbies) // 输出: Hobbies: [reading travelling]

}
