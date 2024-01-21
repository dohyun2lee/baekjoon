package main

import (
	// "bufio"
	//"os"
	//"bufio"
	"fmt"
	// "reflect"
	// 	"strconv"
	// 	"sort"
	"strings"
)

func main() {
	//var a []string
	s := "So when I die (the [first] I will see in (heaven) is a score list)."
	S := strings.Split(s, "")
	
	fmt.Println(S[2])

	if S[2] == " " {
		fmt.Println("SDD")
	}

	// for i:=0;i<len(S);i++ {
	// 	fmt.Println(S[i])
	// }
}
