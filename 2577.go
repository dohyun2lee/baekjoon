package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var A, B, C, mul int
	var zero, one, two, thr, fou, fiv, six, sev, eig, nin int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)
	fmt.Fscanln(reader, &C)
	mul = A * B * C

	ans := strconv.Itoa(mul)

	for i := 0; i < len(ans); i++ {
		switch ans[i] {
		case 48:
			zero++
		case 49:
			one++
		case 50:
			two++
		case 51:
			thr++
		case 52:
			fou++
		case 53:
			fiv++
		case 54:
			six++
		case 55:
			sev++
		case 56:
			eig++
		case 57:
			nin++
		}
	}

	fmt.Println(zero)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(thr)
	fmt.Println(fou)
	fmt.Println(fiv)
	fmt.Println(six)
	fmt.Println(sev)
	fmt.Println(eig)
	fmt.Println(nin)
}
