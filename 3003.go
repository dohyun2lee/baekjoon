	package main

	import (
		"fmt"
	)

	func main() {
		var king, queen, rook, bishop, knight, pawn int

		fmt.Scanln(&king, &queen, &rook, &bishop, &knight, &pawn)

		fmt.Println(1-king,1-queen, 2-rook, 2-bishop, 2-knight, 8-pawn)
	}