package main

import (
	// "github.com/gin-gonic/gin"

	"fmt"
	"gomod/reflecttest"
)

func main() {
	fmt.Println("hello world begin ! -------------")
	// mylib.ContainerTest()
	// mylib.TimeTest()
	// ostest.OsTest()
	// filerw.FileRW()

	reflecttest.ReflectTest()
}
