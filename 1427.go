package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	var n string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)
	N := strings.Split(n, "")	

	sort.Strings(N)

	for i:=len(N)-1;i>=0;i-- {
		fmt.Fprintf(writer, "%s", N[i])
	}
}