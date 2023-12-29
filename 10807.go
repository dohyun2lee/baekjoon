package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var a, c, count int

	fmt.Scan(&a)
	var slice []int = make([]int, a)

	stdin := bufio.NewReader(os.Stdin)

	for i:=0;i<a;i++{
		var b int
		_, err:=fmt.Scan(&b)

		if err!=nil{
			fmt.Println(err)
			stdin.ReadString('\n')
		} else {
			slice[i] = b
		}
	}

	fmt.Scan(&c)

	for i:=0;i<a;i++{
		if c == slice[i]{
			count += 1
		}
	}

	fmt.Println(count)
}