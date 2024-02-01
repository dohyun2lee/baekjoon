package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
)

func main() {
	var T int
	var x1, x2, y1, y2, r1, r2 float64
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &x1, &y1, &r1, &x2, &y2, &r2)

		between := math.Sqrt(math.Pow(x1 - x2, 2)+ math.Pow(y1 - y2, 2))
		
		if between == (r1 + r2) && between != 0 {
			fmt.Fprintln(writer, 1)
		} else if between == math.Abs(r1 - r2) && between != 0 {
			fmt.Fprintln(writer, 1)
		} else if math.Abs(r1 - r2) < between && between < r1 + r2 {
			fmt.Fprintln(writer, 2)
		} else if between < math.Abs(r1 - r2) {
			fmt.Fprintln(writer, 0)
		} else if between > (r1 + r2) {
			fmt.Fprintln(writer, 0)
		} else if between == 0 {
			if r1 == r2 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		}
	}
}