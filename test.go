package main

import (
	// "bufio"
	// "os"
	"fmt"
	// "reflect"
	// 	"strconv"
	// 	"sort"
 	// "strings"
 )

func main() {
	var a []int
	var x int
	
	for i:=0;i<3;i++ {
		fmt.Scanln(&x)
		a = append(a, x)
	}

	fmt.Println(a)
	a = a[:len(a)-1]
	fmt.Println(a)
}
