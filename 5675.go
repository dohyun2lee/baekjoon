package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var time = []int{0,132,6,138,12,144,18,150,24,156,30,162,36,168,42,174,48,180,54,60,66,72,78,84,90,96,102,108,114,120,126}
	var A int
	var angle map[int]bool = make(map[int]bool)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i:=0;i<len(time);i++ {
		angle[time[i]] = true
	}

	for {
		cnt, _ := fmt.Fscanln(reader, &A)

		if cnt != 1 {
			break
		}

		if angle[A] {
			fmt.Fprintln(writer, "Y")
		} else {
			fmt.Fprintln(writer, "N")
		}
	}
}
