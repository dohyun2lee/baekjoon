package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	var A, B string
	var A_5 []string
	var A_6 []string
	var B_5 []string
	var B_6 []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B)

	a := strings.Split(A, "")
	b := strings.Split(B, "")

	for i:=0;i<len(a);i++ {
		if a[i] == "5" || a[i] == "6" {
			A_5 = append(A_5, "5")
			A_6 = append(A_6, "6")
		} else {
			A_5 = append(A_5, a[i])
			A_6 = append(A_6, a[i])
		}
	}
	for i:=0;i<len(b);i++ {
		if b[i] == "5" || b[i] == "6" {
			B_5 = append(B_5, "5")
			B_6 = append(B_6, "6")
		} else {
			B_5 = append(B_5, b[i])
			B_6 = append(B_6, b[i])
		}
	}

	A5, _ := strconv.Atoi(strings.Join(A_5, ""))
	A6, _ := strconv.Atoi(strings.Join(A_6, ""))
	B5, _ := strconv.Atoi(strings.Join(B_5, ""))
	B6, _ := strconv.Atoi(strings.Join(B_6, ""))

	fmt.Printf("%d %d\n",A5+B5, A6+B6)
}