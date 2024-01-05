package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var S string
	reader := bufio.NewReader(os.Stdin)
	var cnt int
	var slice []string

	fmt.Fscanln(reader, &S)

	s := strings.Split(S, "")

	for i:=0;i<len(s);i++ {
		slice = append(slice, s[i])
	}

	slice = append(slice, "")
	slice = append(slice, "")

	for i:=0;i<len(slice);i++ {		
		if slice[i] == "" {
			break
		}
		if slice[i] == "c" && slice[i+1] == "=" {
			cnt++
			i++
		} else if slice[i] == "c" && slice[i+1] == "-" {
			cnt++
			i++
		} else if slice[i] == "d" && slice[i+1] == "z" && slice[i+2] == "=" {
			cnt++
			i += 2
		} else if slice[i] == "d" && slice[i+1] == "-" {
			cnt++
			i++
		} else if slice[i] == "l" && slice[i+1] == "j" {
			cnt++
			i++
		} else if slice[i] == "n" && slice[i+1] == "j" {
			cnt++
			i++
		} else if slice[i] == "s" && slice[i+1] == "=" {
			cnt++
			i++
		} else if slice[i] == "z" && slice[i+1] == "=" {
			cnt++
			i++
		} else {
			cnt++
		}
	}
	
	fmt.Println(cnt)
}