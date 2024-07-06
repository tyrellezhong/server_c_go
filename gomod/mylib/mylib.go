package mylib

import (
	"container/list"
	"fmt"
	"sort"
)

func Sum(a, b int) int {
	return a + b
}

func ContainerTest() {
	SliceTest()
	MapTest()
	ListTest()
}

func SliceTest() {
	fmt.Println("SliceTest ------------------------------")
	var slice = make([]int, 5, 50)
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3, 9, 8, 7)
	slice = slice[5:]
	sort.Ints(slice)

	for i, val := range slice {
		fmt.Println(i, val)
	}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(i, slice[i])
	// }
}

func MapTest() {
	fmt.Println("MapTest ------------------------------")
	var dic = make(map[int]string)
	dic[1] = "a"
	dic[2] = "b"
	delete(dic, 1)
	if val, ok := dic[2]; ok {
		fmt.Println(val)
	}
}

func ListTest() {
	fmt.Println("ListTest ------------------------------")
	li := list.New()
	li.PushBack("123")
	e := li.PushBack("456")
	li.InsertBefore("---", e)

	for n := li.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}

}
