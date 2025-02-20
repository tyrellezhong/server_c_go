package main

import (
	_ "embed"
	"fmt"
	txttemplate "gomod/txt_template"
)

type BKResult struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
}

func main() {
	fmt.Println("hello world begin ! -------------")

	txttemplate.ConditionTest()
}
