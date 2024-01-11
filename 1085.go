package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var x, y, w, h, width, height int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &x, &y, &w, &h)

	if x < w - x {
		width = x
	} else {
		width = w - x
	}

	if y < h - y {
		height = y
	} else {
		height = h - y
	}

	if width < height {
		fmt.Println(width)
	} else {
		fmt.Println(height)
	}
}