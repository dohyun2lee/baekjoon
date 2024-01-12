package main

import "fmt"

func main() {
	var a, b, c, d, e, f, x, y int

	fmt.Scanln(&a, &b)
	fmt.Scanln(&c, &d)
	fmt.Scanln(&e, &f)

	if a == c{
		x = e
	} else if a == e{
		x = c
	} else {
		x = a
	}
	if b == d{
		y = f
	} else if b == f{
		y = d
	} else {
		y = b
	}

	fmt.Println(x, y)
}