package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var month, day int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &month, &day)

	switch month {
	case 1:
		day += 0
	case 2:
		day += 3
	case 3:
		day += 3
	case 4:
		day += 6
	case 5:
		day += 1
	case 6:
		day += 4
	case 7:
		day += 6
	case 8:
		day += 2
	case 9:
		day += 5
	case 10:
		day += 0
	case 11:
		day += 3
	case 12:
		day += 5
	}

	day--
	day = day % 7

	switch day {
	case 0:
		fmt.Fprintln(writer, "MON")
	case 1:
		fmt.Fprintln(writer, "TUE")
	case 2:
		fmt.Fprintln(writer, "WED")
	case 3:
		fmt.Fprintln(writer, "THU")
	case 4:
		fmt.Fprintln(writer, "FRI")
	case 5:
		fmt.Fprintln(writer, "SAT")
	case 6:
		fmt.Fprintln(writer, "SUN")
	}
}
