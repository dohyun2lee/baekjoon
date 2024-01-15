package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d, e, f, x, y int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &a, &b, &c, &d, &e, &f)

	 if a == 0 {
		y = c / b
		x = (f - (e * c / b)) / d
	} else if b == 0 {
		x = c / a
		y = (f - (d * c / a)) / e
	} else if d == 0 {
		y = f / e
		x = (c - (b * f / e)) / a
	} else if e == 0{
		x = f / d
		y = (c - (a * f / d)) / b
	} else {
		y = ((c * d) - (a * f)) / ((b * d) - (a * e))
		x = (c - (b * y)) / a
	}

	fmt.Println(x, y)
}
