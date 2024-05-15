package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, cnt int
	var before, after string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	fmt.Fscanln(reader, &before)
	fmt.Fscanln(reader, &after)

	if N % 2 == 0 {
		for i:=0;i<len(before);i++ {
			if before[i] == after[i] {
				cnt++
			} else {
				break
			}
		}
	} else {
		for i:=0;i<len(before);i++ {
			if before[i] != after[i] {
				cnt++
			} else {
				break
			}
		}
	}

	if cnt != len(before) {
		fmt.Fprintln(writer, "Deletion failed")
	} else {
		fmt.Fprintln(writer, "Deletion succeeded")
	}
}