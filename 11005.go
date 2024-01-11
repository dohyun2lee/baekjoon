package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, B int
	var dig []int
	var ans []string
	var res string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &B)

	if N != 0 {
		for N > 0 {
			if (N % B) > 0 {
				dig = append(dig, N%B)
				N = N / B
			} else {
				dig = append(dig, 0)
				N = N / B
			}
		}

		for i := 0; i < len(dig); i++ {
			if dig[i] < 10 {
				ans = append(ans, strconv.Itoa(dig[i]))
			} else {
				switch dig[i] {
				case 10:
					ans = append(ans, "A")
				case 11:
					ans = append(ans, "B")
				case 12:
					ans = append(ans, "C")
				case 13:
					ans = append(ans, "D")
				case 14:
					ans = append(ans, "E")
				case 15:
					ans = append(ans, "F")
				case 16:
					ans = append(ans, "G")
				case 17:
					ans = append(ans, "H")
				case 18:
					ans = append(ans, "I")
				case 19:
					ans = append(ans, "J")
				case 20:
					ans = append(ans, "K")
				case 21:
					ans = append(ans, "L")
				case 22:
					ans = append(ans, "M")
				case 23:
					ans = append(ans, "N")
				case 24:
					ans = append(ans, "O")
				case 25:
					ans = append(ans, "P")
				case 26:
					ans = append(ans, "Q")
				case 27:
					ans = append(ans, "R")
				case 28:
					ans = append(ans, "S")
				case 29:
					ans = append(ans, "T")
				case 30:
					ans = append(ans, "U")
				case 31:
					ans = append(ans, "V")
				case 32:
					ans = append(ans, "W")
				case 33:
					ans = append(ans, "X")
				case 34:
					ans = append(ans, "Y")
				case 35:
					ans = append(ans, "Z")
				}
			}
		}

		for i := 0; i < len(rotate(ans)); i++ {
			res += rotate(ans)[i]
		}

		fmt.Println(res)
	} else {
		fmt.Println("0")
	}
}

func rotate(a []string) []string {
	result := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[len(a)-(i+1)]
	}
	return result
}
