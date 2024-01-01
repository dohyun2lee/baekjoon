package main 

import "fmt"

func main() {
	var N int

	fmt.Scanln(&N)

	for i:=1;i<=N;i++ {
		for j:=0;j<i;j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}