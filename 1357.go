package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var A, B string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &A, &B)

	revA, _ := strconv.Atoi(Rev(A))
	revB, _ := strconv.Atoi(Rev(B))

	ans,_ := strconv.Atoi(Rev(strconv.Itoa(revA + revB)))

	fmt.Fprintln(writer, ans)
}

func Rev(a string) string {
	var res string
	for i := len(a) - 1; i >= 0; i-- {
		res += string(a[i])
	}

	return res
}
