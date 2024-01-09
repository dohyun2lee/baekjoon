package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"math"
)

func main() {
	var N string
	var dig []float64 
	var B, ans float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &B)
	n := strings.Split(N, "")
	

	for i:=0;i<len(n);i++ {
		switch n[i] {
		case "0" :
			dig = append(dig, 0)
		case "1" :
			dig = append(dig, 1)
		case "2" :
			dig = append(dig, 2)
		case "3" :
			dig = append(dig, 3)
		case "4" :
			dig = append(dig, 4)
		case "5" :
			dig = append(dig, 5)
		case "6" :
			dig = append(dig, 6)
		case "7" :
			dig = append(dig, 7)
		case "8" :
			dig = append(dig, 8)
		case "9" :
			dig = append(dig, 9)
		case "A" :
			dig = append(dig, 10)	
		case "B" :
			dig = append(dig, 11)
		case "C" :
			dig = append(dig, 12)	
		case "D" :
			dig = append(dig, 13)
		case "E" :
			dig = append(dig, 14)	
		case "F" :
			dig = append(dig, 15)
		case "G" :
			dig = append(dig, 16)	
		case "H" :
			dig = append(dig, 17)
		case "I" :
			dig = append(dig, 18)	
		case "J" :
			dig = append(dig, 19)
		case "K" :
			dig = append(dig, 20)	
		case "L" :
			dig = append(dig, 21)
		case "M" :
			dig = append(dig, 22)	
		case "N" :
			dig = append(dig, 23)
		case "O" :
			dig = append(dig, 24)	
		case "P" :
			dig = append(dig, 25)
		case "Q" :
			dig = append(dig, 26)	
		case "R" :
			dig = append(dig, 27)
		case "S" :
			dig = append(dig, 28)	
		case "T" :
			dig = append(dig, 29)
		case "U" :
			dig = append(dig, 30)	
		case "V" :
			dig = append(dig, 31)
		case "W" :
			dig = append(dig, 32)	
		case "X" :
			dig = append(dig, 33)
		case "Y" :
			dig = append(dig, 34)	
		case "Z" :
			dig = append(dig, 35)
		}
	}

	for i:=0;i<len(dig);i++ {
		ans += dig[i] * math.Pow(B, float64(len(dig) - i -1))
	}

	fmt.Printf("%d", int(ans))
}