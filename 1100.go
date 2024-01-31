package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	var s string
	var cnt int
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<8;i++ {
		fmt.Fscanln(reader, &s)
		chess := strings.Split(s, "")

		if i%2 == 0 {
			for j:=0;j<8;j+=2 {
				if chess[j] == "F" {
					cnt++
				}
			}
		} else {
			for j:=1;j<8;j+=2 {
				if chess[j] == "F" {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}