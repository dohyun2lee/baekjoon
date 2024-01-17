package main

import (
	"fmt"
	//"reflect"
	//"strconv"
)

func main() {
	var a []string
	var n int

	fmt.Scanln(&n)
	
	for i:=0;i<n;i++ {
		var s string
		fmt.Scanln(&s)
		a = append(a, s)
	}

	ans := uniqueSlice(a)

	fmt.Println(ans)
}

func uniqueSlice(s []string) []string {
    var res = map[string]int{}
	var ans []string

	for i:=0;i<len(s);i++ {
		if res[s[i]] == 0 {
			res[s[i]]++
			ans = append(ans, s[i])
		}
	}

	return ans
}