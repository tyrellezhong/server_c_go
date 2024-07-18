package reflecttest

import (
	"fmt"
	"reflect"
	"strings"
)

type Address struct {
	Home        string
	Company     string
	Phone       string
	ApartmentID int
}

type Person struct {
	Name string `json:"name" default:"zzz"`
	Age  int
	Addr Address
}

func testStructFields(s interface{}, fieldName string, expectedValue interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return false
	}

	return reflect.DeepEqual(field.Interface(), expectedValue)
}

func callFunction(fn interface{}, args ...interface{}) []interface{} {
	v := reflect.ValueOf(fn)
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	results := v.Call(in)
	out := make([]interface{}, len(results))
	for i, result := range results {
		out[i] = result.Interface()
	}

	return out
}

func add(a, b int) int {
	return a + b
}

func ReflectTest() {
	type myint int
	var x myint = 42
	t := reflect.TypeOf(x)
	vx := reflect.ValueOf(x)

	fmt.Println("Type:", t, "kind:", t.Kind())
	fmt.Println("Value:", vx, "kind:", vx.Kind())

	v2 := reflect.ValueOf(&x).Elem() // 获取指针的元素

	if v2.CanInt() {
		v2.SetInt(100)
	}

	fmt.Println("Modified Value:", x)

	// ----
	px := Person{Name: "Alice", Age: 30}

	fmt.Println(testStructFields(px, "Name", "Alice")) // true
	fmt.Println(testStructFields(px, "Age", 30))       // true
	fmt.Println(testStructFields(px, "Age", 25))       // false
	fmt.Println(testStructFields(px, "Height", 170))   // false
	// ------
	results := callFunction(add, 1, 2)
	fmt.Println(results[0]) // 3
	// --------

	// 结构体类型测试
	ins := &Person{
		Name: "Alice",
		Age:  30,
		Addr: Address{
			Home:    "New York",
			Company: "NY",
		},
	}
	// 获取结构体实例的反射类型对象
	typeOfPerson := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfPerson.Name(), typeOfPerson.Kind())
	// 取类型的元素
	typeOfPerson = typeOfPerson.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfPerson.Name(), typeOfPerson.Kind())

	for i := 0; i < typeOfPerson.NumField(); i++ {
		fieldType := typeOfPerson.Field(i)
		fmt.Printf("name:%v pkgpath:%v tag:%v Offset:%v Anonymous:%v\n", fieldType.Name, fieldType.PkgPath, fieldType.Tag, fieldType.Offset, fieldType.Anonymous)
	}

	if nameField, ok := typeOfPerson.FieldByName("Name"); ok {
		fmt.Println(nameField.Tag.Get("json"), nameField.Tag.Get("default"))
	}

	indexField := typeOfPerson.FieldByIndex([]int{0})
	fmt.Println(indexField.Tag.Get("json"), indexField.Tag.Get("default"))

	// 自定义匹配函数
	matchFunc := func(name string) bool {
		return strings.HasPrefix(name, "Name")
	}

	// 通过匹配函数查找字段
	if field, ok := typeOfPerson.FieldByNameFunc(matchFunc); ok {
		fmt.Println(field.Name)
	}

	// 结构体类型值修改
	// 创建一个 Person 实例
	p := Person{
		Name: "Alice",
		Age:  30,
		Addr: Address{
			Home:        "123 Main St",
			Company:     "Tech Corp",
			Phone:       "123-456-7890",
			ApartmentID: 101,
		},
	}

	// 获取结构体的反射值
	v := reflect.ValueOf(&p).Elem()

	// 修改 Name 字段
	nameField := v.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Bob")
	}

	// 修改 Age 字段
	ageField := v.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(25)
	}

	// 修改 Addr 字段中的 Home 字段
	addrField := v.FieldByName("Addr")
	if addrField.IsValid() && addrField.CanSet() {
		homeField := addrField.FieldByName("Home")
		if homeField.IsValid() && homeField.CanSet() {
			homeField.SetString("456 Elm St")
		}

		// 修改 Addr 字段中的 Company 字段
		companyField := addrField.FieldByName("Company")
		if companyField.IsValid() && companyField.CanSet() {
			companyField.SetString("New Tech Corp")
		}

		// 修改 Addr 字段中的 Phone 字段
		phoneField := addrField.FieldByName("Phone")
		if phoneField.IsValid() && phoneField.CanSet() {
			phoneField.SetString("987-654-3210")
		}

		// 修改 Addr 字段中的 ApartmentID 字段
		apartmentIDField := addrField.FieldByName("ApartmentID")
		if apartmentIDField.IsValid() && apartmentIDField.CanSet() {
			apartmentIDField.SetInt(202)
		}
	}

	// 打印修改后的结构体
	fmt.Printf("%+v\n", p)
}
