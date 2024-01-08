package main

import (
	"fmt"
)

func main() {
	var one, two, thr int

	fmt.Scanln(&one, &two, &thr)

	if one == two && two == thr {
		fmt.Println(one * 1000 + 10000)
	} else if one == two && two != thr {
		fmt.Println(one * 100 + 1000)
	} else if one == thr && two != thr {
		fmt.Println(one * 100 + 1000)
	} else if two == thr && one != thr {
		fmt.Println(two * 100 + 1000)
	} else {
		if one >= two && one >= thr {
			fmt.Println(one * 100)
		} else if two >= one && two >= thr {
			fmt.Println(two * 100)
		} else {
			fmt.Println(thr * 100)
		}
	}
}