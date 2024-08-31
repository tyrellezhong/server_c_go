package main

import (
	_ "embed"
	"fmt"
	"gomod/gogenerate"
	"gomod/httpex"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

//go:embed service.gotpl
var serviceTemplate string

func main3() {
	// 解析模板
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// 定义模板数据
	data := struct {
		PackageName string
		ServiceName string
	}{
		PackageName: "example",
		ServiceName: "ExampleService",
	}

	// 执行模板并输出到标准输出
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func main2() {
	fmt.Println("hello world begin ! -------------")
	// riskInfo := &PayRickControlInfo{
	// 	AdultStatus:         -1,
	// 	RegisterRegion:      "\"KR\"",
	// 	PlayerAge:           13,
	// 	LastModifyTimestamp: 4,
	// }
	// logger, _ := zap.NewProduction()
	// defer logger.Sync()
	// logger.Info("pay risk control check", zap.Any("riskctl info", riskInfo))

	// mylib.ContainerTest()
	// mylib.TimeTest()
	// ostest.OsTest()
	// filerw.FileRW()

	// // reflecttest.ReflectTest()
	// mylib.TimeParse()

	httpex.UrlEncode()
	httpex.UrlDecode()

	// testmap := make(map[int]int)
	// value := testmap[1]
	// if value == 0 {
	// 	fmt.Println("is zero")
	// }

	days := []gogenerate.Weekday{gogenerate.Sunday, gogenerate.Monday, gogenerate.Tuesday}
	for _, day := range days {
		fmt.Println(day.String())
	}
}

// 定义结构体来映射 YAML 数据
type Config struct {
	App      AppConfig      `yaml:"app"`
	Features []string       `yaml:"features"` // 将 features 定义为切片
	Database DatabaseConfig `yaml:"database"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Debug   bool   `yaml:"debug"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AlarmControl struct {
	NotAlarmMeshSpace []string `yaml:"not_alarm_mesh_space"`
	NotAlarmZone      []string `yaml:"not_alarm_zone"`
	MatchAlarmEnv     []string `yaml:"match_alarm_envi"`
	NotAlarmContent   []string `yaml:"not_alarm_content"`
	Rtx               []string `yaml:"rtx"`
	User              string   `yaml:"user"`
	SliceTest         []string `yaml:"slice_test"`
}

func main() {
	var err error
	// 读取 YAML 文件
	data, err := os.ReadFile("./alarm_config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// // 解析 YAML 数据
	// var config Config
	// err = yaml.Unmarshal(data, &config)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	var alarmCtl AlarmControl
	err = yaml.Unmarshal(data, &alarmCtl)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 打印解析后的数据
	// fmt.Printf("App Name: %s\n", config.App.Name)
	// fmt.Printf("App Version: %s\n", config.App.Version)
	// fmt.Printf("Debug Mode: %t\n", config.App.Debug)
	// fmt.Printf("Features: %v\n", config.Features) // 打印切片
	// fmt.Printf("Database Host: %s\n", config.Database.Host)
	// fmt.Printf("Database Port: %d\n", config.Database.Port)
	// fmt.Printf("Database User: %s\n", config.Database.User)

}
