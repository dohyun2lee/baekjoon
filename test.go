package main 

import (
"fmt"
"strconv"
)

func main() {
	var a string = "abcd"
	var b string = "5"

	b = b + "0"

	x,_ := strconv.Atoi(b)

	fmt.Println(x)
	fmt.Println(a)
}