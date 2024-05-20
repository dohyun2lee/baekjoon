package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, "NFC West       W   L  T")
	fmt.Fprintln(writer, "-----------------------")
	fmt.Fprintln(writer, "Seattle        13  3  0")
	fmt.Fprintln(writer, "San Francisco  12  4  0")
	fmt.Fprintln(writer, "Arizona        10  6  0")
	fmt.Fprintln(writer, "St. Louis      7   9  0")
	fmt.Fprintln(writer, "")
	fmt.Fprintln(writer, "NFC North      W   L  T")
	fmt.Fprintln(writer, "-----------------------")
	fmt.Fprintln(writer, "Green Bay      8   7  1")
	fmt.Fprintln(writer, "Chicago        8   8  0")
	fmt.Fprintln(writer, "Detroit        7   9  0")
	fmt.Fprintln(writer, "Minnesota      5  10  1")
}
