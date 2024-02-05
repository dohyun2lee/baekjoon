package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	printer := bufio.NewWriter(os.Stdout)
	fmt.Fscan(reader, &input)

	if len(input) == 1 && input[0] == '0' {
		fmt.Println(0)
		return
	}

	for i := 0; i < len(input); i++ {
		if i == 0 {
			switch input[i] - '0' {
			case 0: // do nothing
			case 1:
				printer.Write([]byte("1"))
			case 2:
				printer.Write([]byte("10"))
			case 3:
				printer.Write([]byte("11"))
			case 4:
				printer.Write([]byte("100"))
			case 5:
				printer.Write([]byte("101"))
			case 6:
				printer.Write([]byte("110"))
			case 7:
				printer.Write([]byte("111"))
			}
		} else {
			switch input[i] - '0' {
			case 0:
				printer.Write([]byte("000"))
			case 1:
				printer.Write([]byte("001"))
			case 2:
				printer.Write([]byte("010"))
			case 3:
				printer.Write([]byte("011"))
			case 4:
				printer.Write([]byte("100"))
			case 5:
				printer.Write([]byte("101"))
			case 6:
				printer.Write([]byte("110"))
			case 7:
				printer.Write([]byte("111"))
			}
		}
	}
	printer.Flush()
}