package main

import (
    "fmt"
	"bufio"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	
	var i int = 3
	var cnt int = 1

	if N <= 3 {
		fmt.Fprintln(writer,"1")
	} else {
		for {
			N -= i
			if N <= 0 {
				break
			}
			i += 2
			cnt++
		}
		fmt.Println(cnt)
	}
}
