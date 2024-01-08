	package main

	import (
		"bufio"
		"fmt"
		"os"
	)

	func main() {
		var N, M int
		reader := bufio.NewReader(os.Stdin)

		_, err := fmt.Fscanln(reader, &N, &M)

		var matrix_one [101][101]int
		var matrix_two [101][101]int
		var matrix_res [101][101]int

		if err != nil {
			fmt.Println(err)
			reader.ReadString('\n')
		} else {
			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					fmt.Scan(&matrix_one[i][j])
				}
			}
			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					fmt.Scan(&matrix_two[i][j])
				}
			}
			for i := 0; i < N; i++ {
				for j := 0; j <
					M; j++ {
					matrix_res[i][j] = matrix_one[i][j] + matrix_two[i][j]
					fmt.Print(matrix_res[i][j], " ")
				}
				fmt.Println()
			}
		}
	}
