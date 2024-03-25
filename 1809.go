package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

fmt.Fprintln(writer, "(___)")
fmt.Fprintln(writer, "(o o)____/")
fmt.Fprintln(writer,  "@@      \\")
fmt.Fprintln(writer,  " \\ ____,/")
fmt.Fprintln(writer,  " //   //")
fmt.Fprintln(writer,  "^^   ^^")
}