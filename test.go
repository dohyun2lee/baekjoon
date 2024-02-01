package main

import (
	"fmt"
)

func main(){

	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	Kim := 7
	Yim := 11

	var b []int
	for i:=0;i<len(a);i++ {
		if a[i]%2 != 0 {
			b = append(b, a[i]/2+1)
		} else {
			b = append(b, a[i]/2)
		}
	}

	fmt.Println(b)

	Kim = b[Kim-1]
	Yim = b[Yim-1]
	
	fmt.Println(Kim, Yim)

	b = makeSliceUnique(b)
	fmt.Println(b)
}

func makeSliceUnique(x []int) []int {
    keys := make(map[int]struct{}) 
    res := make([]int, 0) 
    for _, val := range x { 
        if _, ok := keys[val]; ok { 
            continue 
        } else { 
            keys[val] = struct{}{} 
            res = append(res, val) 
        } 
    }
    return res 
}