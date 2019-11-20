package main

import "fmt"

func strStr(haystack string, needle string) int {
	var lh = len(haystack)
	var ln = len(needle)
	if lh == 0 && ln == 0 {
		return 0
	} else if lh == 0 {
		return -1
	} else if ln == 0 {
		return 0
	}
	for i := 0; i <= lh-ln; i++ {
		if haystack[i] == needle[0] {
			for j, index := 1, i+1; j < ln; j, index = j+1, index+1 {
				fmt.Println("j----", j)
				fmt.Println("index", index)
				if needle[j] != haystack[index] {
					fmt.Println("break")
					break
				} else if j == ln-1 {
					return i
				}
			}
		}
	}
	return -1
}
func main() {
	res := strStr("mississippi", "issip")
	fmt.Println(res)
}
