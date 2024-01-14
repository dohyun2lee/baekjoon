package main

import "fmt"

func main() {
	var N, res int

	fmt.Scanln(&N)

	for i:=1;i<=N;i++ {
		res = i/1000000 + i%1000000/100000 + i%100000/10000 + i%10000/1000 + i%1000/100 + i%100/10 + i%10 + i
		if res == N {
			fmt.Println(i)
			break
		}
	}

	if res != N {
		fmt.Println("0")
	}
}