package main

import "fmt"

func main() {
	var a, b,c int

	for {
		fmt.Scanln(&a, &b, &c)

		if a == 0 && b == 0 && c==0 {
			break
		} else if a >= b+c || b >= a+c || c >= a+b {
			fmt.Println("Invalid")
		} else if a==b && b==c {
			fmt.Println("Equilateral");
		} else if a != b && b != c && c != a {
			fmt.Println("Scalene");
		} else {
			fmt.Println("Isosceles");
		}
	}
}