package main

import (
	"fmt"
)

func main(){
	var a string

	fmt.Scanln(&a)

	fmt.Printf("%c", a[0])
	fmt.Println(len(a))
}
