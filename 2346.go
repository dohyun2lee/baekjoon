package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, popN, moveNum int
	var popCheck []int
	var location = 1
	var isLeft bool
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscan(reader, &n)

	popCheck = make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &popCheck[i])
	}

	fmt.Println("location :", location)
	fmt.Println("popN :", popN)
	fmt.Println("popcheck :", popCheck)
	fmt.Println("popcheck[location] :", popCheck[location])
	fmt.Println("isLeft :", isLeft)
	fmt.Println("moveNum :", moveNum)
	fmt.Println("--------------------------")
	fmt.Println("--------------------------")

	for {
		moveNum = popCheck[location]
		if moveNum >= 0 {
			isLeft = false
		} else {
			isLeft = true
			moveNum *= -1
		}

		fmt.Println("moveNum :", moveNum)
		fmt.Println("isLeft :", isLeft)
		fmt.Println("--------------------------")

		fmt.Fprintf(writer, "%d ", location)
		popCheck[location] = 0
		popN++

		fmt.Println("popcheck :", popCheck)
		fmt.Println("popN :", popN)
		fmt.Println("--------------------------")

		if popN == n {
			break
		}

		for {
			if moveNum == 0 {
				break
			}

			if isLeft {
				location = (location - 1) % (n + 1)
				if location == 0 {
					location = n
				}
			} else {
				location = (location + 1) % (n + 1)
				if location == 0 {
					location = 1
				}
			}

			fmt.Println("location :", location)

			if popCheck[location] == 0 {
				continue
			}

			moveNum--

			fmt.Println("moveNum :", moveNum)
			fmt.Println("--------------------------")
		}
		fmt.Println("--------------------------")
	}
}
