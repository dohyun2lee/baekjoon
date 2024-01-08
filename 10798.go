package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var matrix [5][15]string
	var Str string
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<5;i++ {
		for j:=0;j<15;j++ {
			matrix[i][j] = ""
		}
	}

	for i:=0;i<5;i++ {
		fmt.Fscan(reader, &Str)
		str := strings.Split(Str, "")
		for j:=0;j<len(str);j++ {
			matrix[i][j] = str[j]
		}
	}

	for i:=0;i<15;i++ {
		for j:=0;j<5;j++ {
			fmt.Print(matrix[j][i])
		}
	}
}