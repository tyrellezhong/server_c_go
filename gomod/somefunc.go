package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

//go:embed service.gotpl
var ServiceTemplate string

func ParseTemplate() {
	// 解析模板
	tmpl, err := template.New("service").Parse(ServiceTemplate)
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

type AlarmControl struct {
	NotAlarmMeshSpace []string `yaml:"not_alarm_mesh_space"`
	NotAlarmZone      []string `yaml:"not_alarm_zone"`
	MatchAlarmEnv     []string `yaml:"match_alarm_envi"`
	NotAlarmContent   []string `yaml:"not_alarm_content"`
	Rtx               []string `yaml:"rtx"`
	User              string   `yaml:"user"`
	SliceTest         []string `yaml:"slice_test"`
}

func ReadYaml() {
	var err error
	// 读取 YAML 文件
	data, err := os.ReadFile("./alarm_config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var alarmCtl AlarmControl
	err = yaml.Unmarshal(data, &alarmCtl)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
