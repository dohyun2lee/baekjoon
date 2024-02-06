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
	var ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &A, &B)

	a := strings.Split(A, "")
	b := strings.Split(B, "")

	for i:=0;i<len(a);i++ {
		for j:=0;j<len(b);j++ {
			x, _ := strconv.Atoi(a[i])
			y, _ := strconv.Atoi(b[j])
			ans += x * y
		}
	}

	fmt.Println(ans)
}
