package main

import (
	_ "embed"
	"gomod/generics"
)

type BKResult struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
}

func groupAnagrams(strs []string) [][]string {
	dic := make(map[[26]int][]string)
	for _, value := range strs {
		hash := [26]int{}
		for _, char := range value {
			hash[char-'a']++
		}
		// dic[hash] = append(dic[hash], value)
		if arr, ok := dic[hash]; ok {
			arr = append(arr, value)
			dic[hash] = arr
		} else {
			dic[hash] = []string{value}
		}
	}
	ret := [][]string{}
	for _, arr := range dic {
		ret = append(ret, arr)
	}
	return ret
}
func main() {

	var s generics.MyString = "TestMyString"

	generics.SayHi(s)
}
