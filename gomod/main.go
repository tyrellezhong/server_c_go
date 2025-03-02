package main

import (
	_ "embed"
	"fmt"
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
	fmt.Println("hello world begin ! -------------")
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	groupAnagrams(arr)
}
func longestConsecutive(nums []int) int {
	dic := make(map[int]bool)
	ret := 0
	for _, value := range nums {
		dic[value] = true
	}
	// for _, value := range nums {
	//     if _, ok := dic[value -1]; !ok {
	//         curNum := value
	//         curCount := 1
	//         for dic[curNum + 1] {
	//             curCount++
	//             curNum++
	//         }
	//         if curCount > ret {
	//             ret = curCount
	//         }
	//     }
	// }
	for _, num := range nums {
		if !dic[num-1] {
			currentNum := num
			curretnStreak := 1
			for dic[currentNum+1] {
				currentNum++
				curretnStreak++
			}
			if curretnStreak > ret {
				ret = curretnStreak
			}
		}
	}
	return ret
}
