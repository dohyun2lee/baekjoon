package main

import "fmt"

func main() {
	var angle1, angle2, angle3 int

	fmt.Scanln(&angle1)
	fmt.Scanln(&angle2)
	fmt.Scanln(&angle3)

	if angle1 + angle2 + angle3 != 180 {
		fmt.Println("Error")
	} else {
		if angle1 == angle2 && angle1 == angle3 {
			fmt.Println("Equilateral")
		} else if angle1 == angle2 {
			fmt.Println("Isosceles")
		} else if angle1 == angle3 {
			fmt.Println("Isosceles") 
		} else if angle2 == angle3 {
			fmt.Println("Isosceles")
		} else {
			fmt.Println("Scalene")
		}
	}
}