package main

import "fmt"

func main(){
	var a []int
	a[1] = 1
	a[2] = 2
	for i:=0;i<5;i++ {
		if i%2==0 {
			continue
		}
		fmt.Println(i)
	}
}