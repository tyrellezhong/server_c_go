package interfaceex

import "fmt"

// 定义一个接口
type Stringer interface {
	String() string
}

// 定义一个结构体成员类型
type MyField struct{}

// 为结构体成员类型实现 String() 方法
func (f MyField) String() string {
	return "MyField String()"
}

// 定义一个结构体，包含匿名成员
type MyStruct struct {
	MyField // 匿名成员
}

// 为结构体实现 String() 方法
// 注释掉此方法以查看匿名成员的行为
func (m MyStruct) String() string {
	return "MyStruct String()"
}

// 如果结构体本身没有实现接口的方法，而其匿名成员实现了该方法，结构体可以通过匿名成员“继承”该方法。
// 如果结构体本身实现了接口的方法，则优先调用结构体自身的方法，而不是匿名成员的方法。
func InterfaceTest() {
	// 创建一个 MyStruct 实例
	myStruct := MyStruct{}

	// 将 MyStruct 实例赋值给 Stringer 接口
	var s Stringer = myStruct

	// 调用 String() 方法
	fmt.Println(s.String()) // 输出: MyField String()
}
