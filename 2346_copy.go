package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, num int
	var ans []int
	var balloon []int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i := 1; i <= N; i++ {
		if i < N {
			fmt.Fscanf(reader, "%d ", &num)
		} else {
			fmt.Fscanln(reader, &num)
		}
		balloon = append(balloon, num)
	}

	fmt.Println("ballon :", balloon)
	fmt.Println("-----------")

	var x int

	fmt.Println("x :", x)
	fmt.Println("balloon[x] :", balloon[x])
	fmt.Println("len :", len(balloon))
	fmt.Println("-----------")

	for len(balloon) > 0 {
		ans = append(ans, x+1)
		fmt.Println("ans :", ans)

		num = balloon[x]
		fmt.Println("num :", num)
		
		balloon = append(balloon[:x], balloon[x+1:]...)
		fmt.Println("balloon :", balloon)

		if num > 0 {
			x = x + ((x + num) % len(balloon))
		} else {
			x = x - ((x - num + 1) % len(balloon))
		}
		fmt.Println("x :", x)
		fmt.Println("len :", len(balloon))
	}

	fmt.Println(ans)
}
