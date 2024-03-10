package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, gcd, lcm int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &a, &b)

	if a > b {
		for {
			lcm+=a
			if lcm%b == 0 {
				break
			}
		}
		for i := b; i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				gcd = i
				break
			}
		}
	} else {
		for {
			lcm+=b
			if lcm%a == 0 {
				break
			}
		}
		for i := a; i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				gcd = i
				break
			}
		}
	}

	fmt.Fprintln(writer, gcd)
	fmt.Fprintln(writer, lcm)
}
