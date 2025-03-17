package txttemplate

import (
	"os"
	"text/template"
	"time"
)

const condition = `
{{- $map := dict "key1" "value1" "key2" "value2" -}}
{{- if .IsAdmin}}
Hello, Admin {{.Name}}! Key1: {{index $map "key1"}}
{{- else}}
Hello, User {{.Name}}! Key2: {{index $map "key2"}}
{{- end}}

{{- range .Items}}
- {{.}}
{{- end}}
`

const tmpl = `
{{define "header"}}
Header
{{end}}

{{define "content"}}
Content: {{.}}
{{end}}

{{define "footer"}}
Footer
{{end}}

{{define "base"}}
{{template "header"}}
{{template "content" "测试自定义传递内容"}}
{{template "footer"}}
{{end}}
`

type Person struct {
	Name    string
	IsAdmin bool
	Items   []string
}

type Alert struct {
	Status      string
	StartsAt    time.Time
	Labels      map[string]string
	Annotations map[string]string
}

type Data struct {
	Status string
	Alerts struct {
		Firing   []Alert
		Resolved []Alert
	}
	GroupLabels       map[string]string
	CommonLabels      map[string]string
	CommonAnnotations map[string]string
}

// 示例数据
var data = Data{
	Status: "firing",
	Alerts: struct {
		Firing   []Alert
		Resolved []Alert
	}{
		Firing: []Alert{
			{
				Status:      "firing",
				StartsAt:    time.Now(),
				Labels:      map[string]string{"alertname": "HighCPUUsage", "severity": "critical", "instance": "localhost", "cpu": "90%"},
				Annotations: map[string]string{"summary": "High CPU usage", "description": "CPU usage is above 90%"},
			},
			{
				Status:      "firing",
				StartsAt:    time.Now(),
				Labels:      map[string]string{"alertname": "MemoryUsage", "severity": "critical", "instance": "localhost", "memory": "80%"},
				Annotations: map[string]string{"summary": "High CPU usage", "description": "CPU usage is above 90%"},
			},
		},
		Resolved: []Alert{},
	},
	GroupLabels:       map[string]string{"alertname": "HighCPUUsage"},
	CommonLabels:      map[string]string{"alertname": "HighCPUUsage", "severity": "critical"},
	CommonAnnotations: map[string]string{"summary": "High CPU usage"},
}

func ConditionTest() {

	// funcMap := template.FuncMap{
	// 	"dict": func(values ...interface{}) (map[string]interface{}, error) {
	// 		if len(values)%2 != 0 {
	// 			return nil, fmt.Errorf("invalid dict call")
	// 		}
	// 		dict := make(map[string]interface{}, len(values)/2)
	// 		for i := 0; i < len(values); i += 2 {
	// 			key, ok := values[i].(string)
	// 			if !ok {
	// 				return nil, fmt.Errorf("dict keys must be strings")
	// 			}
	// 			dict[key] = values[i+1]
	// 		}
	// 		return dict, nil
	// 	},
	// }

	// t := template.Must(template.New("example").Funcs(funcMap).Parse(condition))
	// p := Person{
	// 	Name:    "Alice",
	// 	IsAdmin: true,
	// 	Items:   []string{"Item1", "Item2", "Item3"},
	// }
	// t.Execute(nil, p)

	// // 解析模板
	// t2 := template.Must(template.New("").Parse(tmpl))

	// // 执行模板
	// t2.ExecuteTemplate(nil, "base", "Hello, World!")

	t3 := template.Must(template.New("fileexamp").ParseFiles("../gomod/txt_template/grafana_noticepoint.tmpl"))

	t3.ExecuteTemplate(os.Stdout, "message_at", "vlog")

}
