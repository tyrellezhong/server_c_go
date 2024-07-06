package main

import "gomod/mylib"

func main() {
	println("hello world")
	println("sum is", mylib.Sum(10, 20))
	mylib.ContainerTest()

}
